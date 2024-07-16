package hash_test

import (
	"reflect"
	"testing"

	hs "github.com/3JoB/ulib/hash"
	"github.com/3JoB/ulib/internal/hash"
)

func TestSHA3_224(t *testing.T) {
	type args struct {
		data []byte
	}
	var tests []struct {
		name string
		args args
		want *hash.Hash
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.SHA3_224(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHA3_224() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA3_256(t *testing.T) {
	type args struct {
		data []byte
	}
	var tests []struct {
		name string
		args args
		want *hash.Hash
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.SHA3_256(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHA3_256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA3_384(t *testing.T) {
	type args struct {
		data []byte
	}
	var tests []struct {
		name string
		args args
		want *hash.Hash
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.SHA3_384(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHA3_384() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA3_512(t *testing.T) {
	type args struct {
		data []byte
	}
	var tests []struct {
		name string
		args args
		want *hash.Hash
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.SHA3_512(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHA3_512() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA224(t *testing.T) {
	type args struct {
		data []byte
	}
	var tests []struct {
		name string
		args args
		want *hash.Hash
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.SHA224(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHA224() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA256(t *testing.T) {
	type args struct {
		data []byte
	}
	var tests []struct {
		name string
		args args
		want *hash.Hash
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.SHA256(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA384(t *testing.T) {
	type args struct {
		data []byte
	}
	var tests []struct {
		name string
		args args
		want *hash.Hash
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.SHA384(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHA384() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA512(t *testing.T) {
	type args struct {
		data []byte
	}
	var tests []struct {
		name string
		args args
		want *hash.Hash
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.SHA512(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHA512() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA512_224(t *testing.T) {
	type args struct {
		data []byte
	}
	var tests []struct {
		name string
		args args
		want *hash.Hash
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.SHA512_224(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHA512_224() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA512_256(t *testing.T) {
	type args struct {
		data []byte
	}
	var tests []struct {
		name string
		args args
		want *hash.Hash
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.SHA512_256(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHA512_256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMD5(t *testing.T) {
	type args struct {
		data []byte
	}
	var tests []struct {
		name string
		args args
		want *hash.Hash
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.MD5(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashBcrypt(t *testing.T) {
	type args struct {
		password string
	}
	var tests []struct {
		name string
		args args
		want []byte
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.HashBcrypt(tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashBcrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCorrectBcrypt(t *testing.T) {
	type args struct {
		hash     []byte
		password string
	}
	var tests []struct {
		name string
		args args
		want bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.CorrectBcrypt(tt.args.hash, tt.args.password); got != tt.want {
				t.Errorf("CorrectBcrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMD5Str(t *testing.T) {
	type args struct {
		data []byte
	}
	var tests []struct {
		name string
		args args
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hs.MD5Str(tt.args.data); got != tt.want {
				t.Errorf("MD5Str() = %v, want %v", got, tt.want)
			}
		})
	}
}
