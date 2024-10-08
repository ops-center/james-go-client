package inbox

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	identityapis "kmodules.xyz/resource-metadata/apis/identity/v1alpha1"
	"kmodules.xyz/resource-metadata/client/clientset/versioned/typed/identity/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	identityClient *v1alpha1.IdentityV1alpha1Client
	once           sync.Once
)

func getInboxServiceToken() (string, error) {
	if token, ok := os.LookupEnv("INBOX_SERVER_TOKEN"); ok && len(token) > 0 {
		return token, nil
	}

	once.Do(func() {
		if config, err := ctrl.GetConfig(); err != nil {
			log.Fatalf("error fetching config, %v", err)
		} else if identityClient, err = v1alpha1.NewForConfig(config); err != nil {
			log.Fatalf("error building clientset, %v", err)
		}
	})

	tokenRequest := identityClient.InboxTokenRequests()
	tokenResp, err := tokenRequest.Create(context.TODO(), &identityapis.InboxTokenRequest{
		TypeMeta: v1.TypeMeta{
			Kind:       "InboxTokenRequest",
			APIVersion: "identity.k8s.appscode.com/v1alpha1",
		},
	}, v1.CreateOptions{})

	if err != nil || tokenResp.Response == nil {
		return "", fmt.Errorf("error creating authtoken %v", err)
	}

	token := tokenResp.Response.AgentJWTToken
	if token == "" {
		return "", fmt.Errorf("invalid agent token: token is empty")
	}

	return token, nil
}
