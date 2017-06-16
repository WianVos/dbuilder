package config

import (
	"reflect"
	"testing"
)

var jsonStream = `{
    "settings":{
        "build_dir":".build",
        "filemode": "0770"
        },
    "environments":{
        "test": {
            "variables" : {
                "dude": "Wian Vos",
                "base_image": "testimage",
                "justtesting": "effegeeninspiratiemeer"
            }
        },
        "default": {
            "variables" : {
                "dude": "dWian Vos",
                "base_image": "dtestimage",
                "justtesting": "deffegeeninspiratiemeer"
            }
        }
    }
}
`

var invalidJsonStream = `{
    "settings":{
        "build_dir": ".build"
        "filemode": "0770"
        },
    "environments":{
        "test": {
            "variables" : {
                "dude": "Wian Vos",
                "base_image": "testimage",
                "justtesting": "effegeeninspiratiemeer"
            }
        },
        "default" {
            "variables" : {
                "dude": "dWian Vos",
                "base_image": "dtestimage",
                "justtesting": "deffegeeninspiratiemeer"
            }
        }
    }
}
`

var testConfig = Config{
	Settings: SettingsType{
		BuildDir: ".build",
		Filemode: "0770",
	},
	Environments: Environments{
		"test": Environment{
			Variables: Variables{
				"dude":        "Wian Vos",
				"base_image":  "testimage",
				"justtesting": "effegeeninspiratiemeer",
			},
		},
		"default": Environment{
			Variables: Variables{
				"dude":        "dWian Vos",
				"base_image":  "dtestimage",
				"justtesting": "deffegeeninspiratiemeer",
			},
		},
	},
}

// func TestNewConfigFromFile(t *testing.T) {
// 	type args struct {
// 		f string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    Config
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := NewConfigFromFile(tt.args.f)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NewConfigFromFile() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewConfigFromFile() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_readConfigFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{name: "test correct execution", args: args{s: jsonStream}, want: testConfig, wantErr: false},
		{name: "test incorrect execution", args: args{s: invalidJsonStream}, want: Config{}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readConfigFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("readConfigFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readConfigFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
