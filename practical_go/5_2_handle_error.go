package main

import (
	"context"
	"fmt"
	"time"
)

type User any

func getInvitedUserWithEmail(ctx context.Context, email string) (*User, error) {
	var user *User
	return user, nil
}

func badHandling() (*User, error) {
	var email string
	user, err := getInvitedUserWithEmail(context.TODO(), email)
	if err != nil {
		// 発生したエラーをそのまま呼び出し元に返却
		return nil, err
	}
	return user, nil
}

func goodHandling() (*User, error) {
	var email string
	user, err := getInvitedUserWithEmail(context.TODO(), email)
	if err != nil {
		// 呼び出し元で発生したエラーをラップし、付加情報を付与して呼び出し元に返却
		return nil, fmt.Errorf("fail to get invited user with email (%s): %w", email, err)
	}
	return user, nil
}

func retry() error {
	var b []byte

	ctx := context.Background()
	err := retry.Exponential(ctx, 1*time.Second, func(ctx context.Context) error {
		_, ierr := tcpClient.Read(b)
		return ierr
	})
	if err != nil {
		return fmt.Errorf("fail to read from tcp client: %w", err)
	}
	return nil
}
