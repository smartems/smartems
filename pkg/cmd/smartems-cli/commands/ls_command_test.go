package commands

import (
	"errors"
	"github.com/smartems/smartems/pkg/cmd/smartems-cli/commands/commandstest"
	s "github.com/smartems/smartems/pkg/cmd/smartems-cli/services"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMissingPath(t *testing.T) {
	var org = validateLsCommand

	Convey("ls command", t, func() {
		validateLsCommand = org

		Convey("Missing path", func() {
			commandLine := &commandstest.FakeCommandLine{
				CliArgs: []string{"ls"},
				GlobalFlags: &commandstest.FakeFlagger{
					Data: map[string]interface{}{
						"path": "",
					},
				},
			}
			s.IoHelper = &commandstest.FakeIoUtil{}

			Convey("should return error", func() {
				err := lsCommand(commandLine)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("Path is not a directory", func() {
			commandLine := &commandstest.FakeCommandLine{
				CliArgs: []string{"ls"},
				GlobalFlags: &commandstest.FakeFlagger{
					Data: map[string]interface{}{
						"path": "/var/lib/smartems/plugins",
					},
				},
			}

			s.IoHelper = &commandstest.FakeIoUtil{
				FakeIsDirectory: false,
			}

			Convey("should return error", func() {
				err := lsCommand(commandLine)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("can override validateLsCommand", func() {
			commandLine := &commandstest.FakeCommandLine{
				CliArgs: []string{"ls"},
				GlobalFlags: &commandstest.FakeFlagger{
					Data: map[string]interface{}{
						"path": "/var/lib/smartems/plugins",
					},
				},
			}

			validateLsCommand = func(pluginDir string) error {
				return errors.New("dummie error")
			}

			Convey("should return error", func() {
				err := lsCommand(commandLine)
				So(err.Error(), ShouldEqual, "dummie error")
			})
		})

		Convey("Validate that validateLsCommand is reset", func() {
			commandLine := &commandstest.FakeCommandLine{
				CliArgs: []string{"ls"},
				GlobalFlags: &commandstest.FakeFlagger{
					Data: map[string]interface{}{
						"path": "/var/lib/smartems/plugins",
					},
				},
			}

			Convey("should return error", func() {
				err := lsCommand(commandLine)
				So(err.Error(), ShouldNotEqual, "dummie error")
			})
		})
	})
}
