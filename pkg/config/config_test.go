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

var defaultTestConfig = Config{
	Settings: SettingsType{
		BuildDir: ".build",
		Filemode: "0770",
	},
	Environments: Environments{
		"test": Environment{
			Variables: Variables{
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

var emptyDefaultTestConfig = Config{
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

// func TestGetVariablesForEnvironment(t *testing.T) {
// 	type args struct {
// 		e string
// 		d bool
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want map[string]string
// 	}{
// 		{
// 			name: "with default",
// 			args: args{e: "test", d: true},
// 			want: map[string]string{
// 				"dude":        "dWian Vos",
// 				"base_image":  "testimage",
// 				"justtesting": "effegeeninspiratiemeer"},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := GetVariablesForEnvironment(tt.args.e, tt.args.d); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetVariablesForEnvironment() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestConfig_GetVariablesForEnvironment(t *testing.T) {
	type fields struct {
		Settings     SettingsType
		Environments Environments
	}
	type args struct {
		e string
		d bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]string
	}{
		{
			name: "with default",
			fields: fields{
				Settings:     defaultTestConfig.Settings,
				Environments: defaultTestConfig.Environments,
			},
			args: args{e: "test", d: true},
			want: map[string]string{
				"dude":        "dWian Vos",
				"base_image":  "testimage",
				"justtesting": "effegeeninspiratiemeer"},
		},
		{
			name: "without default",
			fields: fields{
				Settings:     defaultTestConfig.Settings,
				Environments: defaultTestConfig.Environments,
			},
			args: args{e: "test", d: false},
			want: map[string]string{
				"base_image":  "testimage",
				"justtesting": "effegeeninspiratiemeer"},
		},
		{
			name: "with empty default",
			fields: fields{
				Settings:     emptyDefaultTestConfig.Settings,
				Environments: emptyDefaultTestConfig.Environments,
			},
			args: args{e: "test", d: true},
			want: map[string]string{
				"dude":        "Wian Vos",
				"base_image":  "testimage",
				"justtesting": "effegeeninspiratiemeer"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Settings:     tt.fields.Settings,
				Environments: tt.fields.Environments,
			}
			if got := c.GetVariablesForEnvironment(tt.args.e, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.GetVariablesForEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}
