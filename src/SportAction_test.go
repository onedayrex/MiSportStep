package src

import (
	"log"
	"testing"
)

func TestSport_login(t *testing.T) {
	type fields struct {
		UserName string
		Password string
		StepRang string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "loginFunction",
			fields: fields{
				UserName: "",
				Password: "",
				StepRang: "200-599",
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sport{
				UserName: tt.fields.UserName,
				Password: tt.fields.Password,
				StepRang: tt.fields.StepRang,
			}
			s.Login()
			s.PushSetp()
		})
	}
}

func TestSport_PushSetp(t *testing.T) {
	type fields struct {
		UserName string
		Password string
		StepRang string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "loginFunction",
			fields: fields{
				UserName: "",
				Password: "",
				StepRang: "200-599",
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sport{
				UserName: tt.fields.UserName,
				Password: tt.fields.Password,
				StepRang: tt.fields.StepRang,
			}
			step := s.RandomStep()
			log.Println(step)
		})
	}
}
