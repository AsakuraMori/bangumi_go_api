package bgm_v0

import (
	"context"
	"fmt"
	"github.com/AsakuraMori/bangumi_go_api/pkg/common"
	"testing"
)

func TestClient_GetCharacterById(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetCharacterById(ctx, "", "77")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}

func TestClient_GetCharacterRelatedSubject(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetCharacterRelatedSubject(ctx, "", "77")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}

func TestClient_GetCharacterRelatedPerson(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetCharacterRelatedPerson(ctx, "", "77")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}
