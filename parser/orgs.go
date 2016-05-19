package parser

import (
	"fmt"
	docopt "github.com/docopt/docopt-go"
	"github.com/sjkyspa/stacks/client/cmd"
	"regexp"
)

func Orgs(argv []string) error {
	usage := `
Valid commands for apps:

orgs:create             create a new organization
orgs:info               view info about an organization
orgs:current            set org as a default org
orgs:members            list members of the organization

Use 'cde help [command]' to learn more.
`
	switch argv[0] {
	case "orgs:create":
		return orgCreate(argv)
	case "orgs:info":
		return orgInfo(argv)
	case "orgs:current":
		return orgCurrent(argv)
	case "orgs:members":
		return orgMembers(argv)
	default:
		if printHelp(argv, usage) {
			return nil
		}

		PrintUsage()
		return nil
	}
	return nil
}

func orgCreate(argv []string) error {
	usage := `
Creates a new organization.

Usage: cde orgs:create [options]

Options:
  -o --org=<name>
    tell system to deploy this app or not, 1 means need, 0 mean no, default 1
`
	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	name := safeGetOrDefault(args, "--org", "1")

	regex := regexp.MustCompile(`^[a-z0-9\-]+$`)
	if !regex.MatchString(name) {
		return fmt.Errorf("'%s' does not match the pattern '[a-z0-9-]+'\n", name)
	}

	return cmd.OrgCreate(name)
}

func orgInfo(argv []string) error {
	usage := `
Prints info about the current organization.

Usage: cde orgs:info [options]

Options:
  -o --org=<org>
    the uniquely identifiable id for the organization.
`
	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	orgName := safeGetValue(args, "--org")

	return cmd.GetOrg(orgName)
}

func orgCurrent(argv []string) error {
	usage := `
Set org as a default org.

Usage: cde orgs:current [options]

Options:
  -o --org=<org>
    the uniquely identifiable id for the organization.
`
	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	orgName := safeGetValue(args, "--org")

	return cmd.SetCurrentOrg(orgName)
}

func orgMembers(argv []string) error {
	usage := `
List members of the organization

Usage: cde orgs:members [options]

Options:
  -o --org=<org>
    the uniquely identifiable id for the organization.
`
	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	orgName := safeGetValue(args, "--org")

	return cmd.ListMembers(orgName)
}



