package algorithm

import (
	"complexity/pkg/settings"
	"complexity/utils/assertions"
	testUtils "complexity/utils/test"
	"testing"
)

func TestCountCharacteristicV3StandardWeights(t *testing.T) {
	commonSettings := testUtils.ReadSettings[settings.SimpleSettings](t, "testdata/common-settings.json")

	tests := []struct {
		name string
		args CharacteristicArgs
		want float64
	}{
		{
			name: "no channels test",
			args: GetArgs(t, commonSettings, "testdata/no-channels-net.xml"),
			want: 1.0,
		},
		{
			name: "1 connection 1 channel",
			args: GetArgs(t, commonSettings, "testdata/2-agents-v2.xml"),
			want: 0.916666666666,
		},
		{
			name: "2 agents, 1 channel, 2 connections test",
			args: GetArgs(t, commonSettings, "testdata/2-agents-v3.xml"),
			want: 0.888888888888,
		},
		{
			name: "2 agents, 2 channels, 2 and 2 connections test",
			args: GetArgs(t, commonSettings, "testdata/2-agents-v4.xml"),
			want: 0.6388888888888,
		},
		{
			name: "2 agents, 2 channels, 2 and 2 connections regexp settings test",
			args: GetArgs(t, testUtils.ReadSettings[settings.RegexpSettings](t, "testdata/common-settings-regexp.json"), "testdata/2-agents-v4.xml"),
			want: 0.6388888888888,
		},
		{
			name: "2 agents, 2 channels, 2 and 4 connections test",
			args: GetArgs(t, commonSettings, "testdata/2-agents-v5.xml"),
			want: 0.611111111111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountCharacteristicV3(tt.args.net, tt.args.settings); !assertions.IsCorrect(tt.want, got) {
				t.Errorf("CountCharacteristicV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountCharacteristicV3InternalBehaviorWeights(t *testing.T) {
	commonSettings := testUtils.ReadSettings[settings.SimpleSettings](t, "testdata/internal-behavior-test/common-with-internal-behavior.json")

	tests := []struct {
		name string
		args CharacteristicArgs
		want float64
	}{
		{
			name: "no channels test",
			args: GetArgs(t, commonSettings, "testdata/no-channels-net.xml"),
			want: 1.0,
		},
		{
			name: "1 connection 1 channel",
			args: GetArgs(t, commonSettings, "testdata/2-agents-v2.xml"),
			want: 0.916666666666,
		},
		{
			name: "2 agents, 1 channel, 2 connections test",
			args: GetArgs(t, commonSettings, "testdata/2-agents-v3.xml"),
			want: 0.888888888888,
		},
		{
			name: "2 agents, 2 channels, 2 and 2 connections test",
			args: GetArgs(t, commonSettings, "testdata/2-agents-v4.xml"),
			want: 0.6388888888888,
		},
		{
			name: "2 agents, 2 channels, 2 and 2 connections regexp settings test",
			args: GetArgs(t, testUtils.ReadSettings[settings.RegexpSettings](t, "testdata/common-settings-regexp.json"), "testdata/2-agents-v4.xml"),
			want: 0.6388888888888,
		},
		{
			name: "2 agents, 2 channels, 2 and 4 connections test",
			args: GetArgs(t, commonSettings, "testdata/2-agents-v5.xml"),
			want: 0.611111111111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountCharacteristicV3(tt.args.net, tt.args.settings); !assertions.IsCorrect(tt.want, got) {
				t.Errorf("CountCharacteristicV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountCharacteristicV3ExternalBehaviorWeights(t *testing.T) {
	commonSettings := testUtils.ReadSettings[settings.SimpleSettings](t, "testdata/external-behavior-test/common-with-external-behavior.json")

	tests := []struct {
		name string
		args CharacteristicArgs
		want float64
	}{
		{
			name: "no channels test",
			args: GetArgs(t, commonSettings, "testdata/no-channels-net.xml"),
			want: 1.0,
		},
		{
			name: "1 connection 1 channel",
			args: GetArgs(t, commonSettings, "testdata/2-agents-v2.xml"),
			want: 0.833333333333,
		},
		{
			name: "2 agents, 1 channel, 2 connections test",
			args: GetArgs(t, commonSettings, "testdata/2-agents-v3.xml"),
			want: 0.7777777777777,
		},
		{
			name: "2 agents, 2 channels, 2 and 2 connections test",
			args: GetArgs(t, commonSettings, "testdata/2-agents-v4.xml"),
			want: 0.2777777777777,
		},
		{
			name: "2 agents, 2 channels, 2 and 4 connections test",
			args: GetArgs(t, commonSettings, "testdata/2-agents-v5.xml"),
			want: 0.2222222222222,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountCharacteristicV3(tt.args.net, tt.args.settings); !assertions.IsCorrect(tt.want, got) {
				t.Errorf("CountCharacteristicV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
