package controllers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestDeleteHabit(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteHabit(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DeleteHabit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
