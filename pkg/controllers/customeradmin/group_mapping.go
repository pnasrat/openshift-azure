package customeradmin

import (
	"context"
	"reflect"
	"sort"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest/azure"
	v1 "github.com/openshift/api/user/v1"
	userv1client "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	"github.com/sirupsen/logrus"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/util/azureclient"
	azgraphrbac "github.com/openshift/openshift-azure/pkg/util/azureclient/graphrbac"
	"github.com/openshift/openshift-azure/pkg/util/mail"
)

// fromMSGraphGroup syncs the values from the given aad group into kubeGroup
// it returns a boolean "changed" that tells the call if an update is required
func fromMSGraphGroup(log *logrus.Entry, userV1 userv1client.UserV1Interface, kubeGroup *v1.Group, kubeGroupName string, msGroupMembers []graphrbac.User) (*v1.Group, bool) {
	var g *v1.Group
	if kubeGroup == nil {
		g = &v1.Group{
			ObjectMeta: meta_v1.ObjectMeta{
				Name: kubeGroupName,
			},
			Users: []string{},
		}
	} else {
		g = kubeGroup.DeepCopy()
	}
	// list all ocp users for reconcile process
	ocpUserList, err := userV1.Users().List(meta_v1.ListOptions{})
	if err != nil {
		log.Warnf("failed to list users %s", err)
	}

	g.Users = []string{}
	for _, user := range msGroupMembers {
		g.Users = append(g.Users, reconcileUsers(log, ocpUserList.Items, user))
	}
	sort.Strings(g.Users)
	return g, !reflect.DeepEqual(kubeGroup, g)
}

// mailEqual tests for e-mail address equality (before the @ is case-sensitive,
// after is not)
func mailEqual(s, t string) bool {
	a := strings.SplitN(s, "@", 2)
	sPrefix, sDomain := a[0], a[1]

	a = strings.SplitN(t, "@", 2)
	tPrefix, tDomain := a[0], a[1]

	return sPrefix == tPrefix && strings.EqualFold(sDomain, tDomain)
}

// tryOverride checks a given e-mail address against the user list, and if it
// finds a name in the user list that matches (subject to mailEqual above), it
// uses it
func tryOverride(mail string, ocpUserList []v1.User) (string, bool) {
	for _, usr := range ocpUserList {
		if mailEqual(usr.Name, mail) {
			return usr.Name, true
		}
	}
	return mail, false
}

// reconcileUsers will take an external user reference from AAD
// match it with already sign-ed in users and updates required metadata.
// This code is trying to solve 4 usecases for guest accounts:
// 1. Member owner account - Mail = nil, GiveName = owner@home.com, MailNickname is partial email with owner_home.com#EXT#
// 2. Guest user with prefix live.com#user@guest.com in OCP users but not in AAD
// 2.b Guest user with prefix live.com#user@guest.com in OCP users but not in AAD and Mail is nil
// 3. Guest user with no prefix user@trustedGuest.com
// 4. AAD Mail field matched a cluster user
// 5. AAD UserPrincipalName field matched a cluster user
// 6. Normal user/ other usecases - Default: Mail if set; UserPrincipalName if Mail not set
// To check structure:
// az ad group member list -g 44e69b4e-2e70-42df-bb97-3a890730d7b0
// External links:
// https://stackoverflow.com/questions/35727866/azure-ad-appending-ext-to-userprincipalname
// https://github.com/aspnet/Security/issues/1717
// Internal example:
// https://microsofteur-my.sharepoint.com/:w:/g/personal/b-majude_microsoft_com/EQCfupKCHN9Gi1uEqRAiiuUBnw_MfcDQLyldWEcV6gGzBw?e=bwfW9K
// There is a further wrinkle where we see the case of the e-mail domain differ
// from the openshift login name.  If we spot this we match the existing user
// name's domain name's case.
func reconcileUsers(log *logrus.Entry, ocpUserList []v1.User, AADUser graphrbac.User) string {
	// This tries to handle use-case 1
	if AADUser.GivenName != nil &&
		AADUser.Mail == nil &&
		AADUser.MailNickname != nil {
		if strings.Contains(*AADUser.MailNickname, "#EXT#") {
			s := strings.Replace(*AADUser.MailNickname, "#EXT#", "", 1)
			idx := strings.LastIndex(s, "_")
			email := s[:idx] + "@" + s[idx+1:]
			if mail.Validate(email) && strings.EqualFold(email, *AADUser.GivenName) {
				username, _ := tryOverride(*AADUser.GivenName, ocpUserList)
				return username
			}
		}
	}
	// This is optimistic code, trying to catch use-case 2 and 2.b
	if AADUser.MailNickname != nil && strings.Contains(*AADUser.MailNickname, "#EXT#") {
		var email string
		if AADUser.Mail != nil {
			email = *AADUser.Mail
		} else {
			s := strings.Replace(*AADUser.MailNickname, "#EXT#", "", 1)
			idx := strings.LastIndex(s, "_")
			email = s[:idx] + "@" + s[idx+1:]
		}
		for _, usr := range ocpUserList {
			// if OpenShift user contains # - we need to drop it for checking.
			if strings.Contains(usr.Name, "#") {
				loginEmail := strings.SplitN(usr.Name, "#", 2)[1]
				if mailEqual(loginEmail, email) {
					return usr.Name
				}
			}
		}
	}

	//Case 4 AADUser Mail field matched a cluster username
	if AADUser.Mail != nil {
		username, found := tryOverride(*AADUser.Mail, ocpUserList)
		if found {
			return username
		}
	}

	//Case 5 AADUser UserPrincipalName matched a cluster username, even when Mail is set
	if AADUser.UserPrincipalName != nil {
		username, found := tryOverride(*AADUser.UserPrincipalName, ocpUserList)
		if found {
			return username
		}
	}

	// returning Mail if set is default behaviour if none
	// of the use-cases where matched
	if AADUser.Mail != nil {
		return *AADUser.Mail
	}
	// if Mail is not set, return UserPrincipalName
	return *AADUser.UserPrincipalName
}

func newAADGroupsClient(ctx context.Context, log *logrus.Entry, config api.AADIdentityProvider) (azgraphrbac.GroupsClient, error) {
	graphauthorizer, err := azureclient.NewAuthorizer(config.ClientID, config.Secret, config.TenantID, azure.PublicCloud.GraphEndpoint)
	if err != nil {
		return nil, err
	}

	return azgraphrbac.NewGroupsClient(ctx, log, config.TenantID, graphauthorizer), nil
}

func getAADGroupMembers(gc azgraphrbac.GroupsClient, groupID string) ([]graphrbac.User, error) {
	members, err := gc.GetGroupMembers(context.Background(), groupID)
	if err != nil {
		return nil, err
	}
	var users []graphrbac.User
	for _, member := range members {
		if user, ok := member.AsUser(); ok {
			users = append(users, *user)
		}
	}
	return users, nil
}
