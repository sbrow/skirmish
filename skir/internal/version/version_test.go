// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package version

/*
func Test_runVersion(t *testing.T) {
	type args struct {
		cmd  *base.Command
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"noArgs", args{CmdVersion, []string{}}, "skir version " + Version},
		{"oneArg", args{CmdVersion, []string{"arg"}}, `usage: version
Run 'skir help version' for details.
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stderr.
			stdCopy := os.Stderr
			r, w, err := os.Pipe()
			if err != nil {
				fmt.Fprintln(os.Stdout, err)
				os.Exit(1)
			}
			os.Stderr = w
			outC := make(chan string)
			go func() {
				var buf bytes.Buffer
				_, err := io.Copy(&buf, r)
				r.Close()
				if err != nil {
					fmt.Fprintf(os.Stdout, "testing: copying pipe: %v\n", err)
					os.Exit(1)
				}
				outC <- buf.String()
			}()
			defer func() {
				w.Close()
				os.Stderr = stdCopy
				got := <-outC
				got = strings.Replace(got, "\r", "", -1)
				fmt.Fprintln(os.Stdout, got)
				if got != tt.want {
					t.Errorf("wanted: \n\"%s\"\ngot:\n\"%s\"", tt.want, got)
				}
			}()
			runVersion(tt.args.cmd, tt.args.args)
		})
	}
}
*/
