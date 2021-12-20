// Package cli contains source code for the cli
package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/camelcase"
	"github.com/getkin/kin-openapi/openapi3"
	specs "github.com/m3o/m3o-cli/api"
	"github.com/spf13/cobra"
	"go.m3o.com/client"
)

var usage = `m3o call [service] [endpoint] [request] e.g m3o call helloworld Call '{"name": "Alice"}'`

var rootCmd = &cobra.Command{
	Use:   "m3o",
	Short: "M3O is a next generation cloud",
	Long:  `M3O is a next generation cloud platform. A single place to explore, discover and consume public APIs. See https://m3o.com for more info.`,
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Command line client",
	Long:  `Command line client. Call or stream from any service directly.`,
}

var callCmd = &cobra.Command{
	Use:   "call",
	Short: "Call a service",
	Long:  `Call a service on the M3O platform`,
	Run: func(cmd *cobra.Command, args []string) {
		// check for a token
		token := os.Getenv("M3O_API_TOKEN")
		if len(token) == 0 {
			fmt.Fprintln(os.Stderr, "Missing M3O_API_TOKEN environment variable")
			os.Exit(1)
		}

		// get the args
		if len(args) != 3 {
			fmt.Fprintln(os.Stderr, usage)
			os.Exit(1)
		}

		service := args[0]
		endpoint := args[1]
		request := args[2]

		var req map[string]interface{}

		if err := json.Unmarshal([]byte(request), &req); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse request as JSON: %v\n", err)
			os.Exit(1)
		}

		c := client.NewClient(nil)
		c.SetToken(token)

		var rsp json.RawMessage

		if err := c.Call(service, endpoint, req, &rsp); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		var out bytes.Buffer
		defer out.Reset()

		if err := json.Indent(&out, rsp, "", "\t"); err != nil {
			fmt.Fprintln(os.Stderr, "Error while trying to format the response", err)
		}

		fmt.Printf("%v\n", string(out.Bytes()))
	},
}

var exploreCmd = &cobra.Command{
	Use:   "explore",
	Short: "Explore M3O services",
	Long:  `List and search for M3O services`,
}

var exploreListCmd = &cobra.Command{
	Use:   "list",
	Short: "List M3O services",
	Long:  `List available M3O services`,
	Run: func(cmd *cobra.Command, args []string) {
		rsp, err := http.Get("https://api.m3o.com/publicapi/explore/index")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error listing services: %v\n", err)
			os.Exit(1)
		}
		defer rsp.Body.Close()

		b, _ := ioutil.ReadAll(rsp.Body)
		if rsp.StatusCode != 200 {
			fmt.Fprintf(os.Stderr, "Error listing services: %v\n", string(b))
			os.Exit(1)
		}

		exp := new(IndexResponse)

		if err := json.Unmarshal(b, &exp); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse response as JSON: %v\n", err)
			os.Exit(1)
		}

		// print header
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

		fmt.Fprintln(w, "Service\tDescription\tCategory\tURL\tIcon")
		fmt.Fprintln(w, "-------\t-----------\t--------\t---\t----")

		for _, api := range exp.Apis {
			desc := strings.Split(api.Description, "\n")
			fmt.Fprintln(w, strings.TrimSpace(api.DisplayName), "\t", strings.TrimSpace(desc[0]), "\t", api.Category, "\t", "https://m3o.com/"+api.Name, "\t", api.Icon)
		}

		w.Flush()
	},
}

var exploreSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for M3O services",
	Long:  `Search available M3O services`,
	Run: func(cmd *cobra.Command, args []string) {
		rsp, err := http.Get("https://api.m3o.com/publicapi/explore/Search?search_term=" + Query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error searching services: %v\n", err)
			os.Exit(1)
		}
		defer rsp.Body.Close()

		b, _ := ioutil.ReadAll(rsp.Body)
		if rsp.StatusCode != 200 {
			fmt.Fprintf(os.Stderr, "Error searching services: %v\n", string(b))
			os.Exit(1)
		}

		exp := new(SearchResponse)

		if err := json.Unmarshal(b, &exp); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse response as JSON: %v\n", err)
			os.Exit(1)
		}

		// print header
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

		fmt.Fprintln(w, "Service\tDescription\tCategory\tURL\tIcon")
		fmt.Fprintln(w, "-------\t-----------\t--------\t---\t----")

		for _, api := range exp.Apis {
			desc := strings.Split(api.Description, "\n")
			fmt.Fprintln(w, strings.TrimSpace(api.DisplayName), "\t", strings.TrimSpace(desc[0]), "\t", api.Category, "\t", "https://m3o.com/"+api.Name, "\t", api.Icon)
		}

		w.Flush()
	},
}

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Stream a response from a service",
	Long:  `Stream a response from a service on the M3O platform`,
	Run: func(cmd *cobra.Command, args []string) {
		// check for a token
		token := os.Getenv("M3O_API_TOKEN")
		if len(token) == 0 {
			fmt.Fprintln(os.Stderr, "Missing M3O_API_TOKEN environment variable")
			os.Exit(1)
		}

		// get the args
		if len(args) != 3 {
			fmt.Fprintln(os.Stderr, usage)
			os.Exit(1)
		}

		service := args[0]
		endpoint := args[1]
		request := args[2]

		var req map[string]interface{}

		if err := json.Unmarshal([]byte(request), &req); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse request as JSON: %v\n", err)
			os.Exit(1)
		}

		c := client.NewClient(nil)
		c.SetToken(token)

		stream, err := c.Stream(service, endpoint, req)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for {
			var rsp json.RawMessage
			if err := stream.Recv(&rsp); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			var out bytes.Buffer
			defer out.Reset()

			if err := json.Indent(&out, rsp, "", "\t"); err != nil {
				fmt.Fprintln(os.Stderr, "Error while trying to format the response", err)
			}

			fmt.Printf("%v\n", string(out.Bytes()))
		}
	},
}

var (
	Query string
)

// generate a list of commands from the public cli
func generateCLI() []*cobra.Command {
	rsp, err := http.Get("https://api.m3o.com/publicapi/explore/index")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing services: %v\n", err)
		os.Exit(1)
	}
	defer rsp.Body.Close()

	b, _ := ioutil.ReadAll(rsp.Body)
	if rsp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Error listing services: %v\n", string(b))
		os.Exit(1)
	}

	resp := new(IndexResponse)

	if err := json.Unmarshal(b, &resp); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse response as JSON: %v\n", err)
		os.Exit(1)
	}

	var commands []*cobra.Command

	for _, api := range resp.Apis {
		name := api.Name
		description := strings.Split(api.Description, "\n")[0]
		//category := api.Category

		command := &cobra.Command{
			Use:   name,
			Short: description,
			Long:  api.Description,
		}

		var spec *openapi3.T

		// try to pull the api spec from the embedded file
		file, err := specs.Specs.ReadFile("spec/" + name + ".json")
		if err == nil {
			json.Unmarshal(file, &spec)
		}

		// if we can't find the file we need to fetch the def
		if file == nil {
			// TODO: cache this data locally
			rsp, err := http.Get("https://api.m3o.com/publicapi/explore/API?name=" + name)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error getting service %s: %v\n", name, err)
				os.Exit(1)
			}
			defer rsp.Body.Close()

			b, _ := ioutil.ReadAll(rsp.Body)
			if rsp.StatusCode != 200 {
				fmt.Fprintf(os.Stderr, "Error getting service %s: %v\n", name, string(b))
				os.Exit(1)
			}

			if err := json.Unmarshal(b, &spec); err != nil {
				fmt.Printf("Error decoding %s: %v\n", name, err)
				os.Exit(1)
			}
		}

		for _, endpoint := range api.Endpoints {
			parts := strings.Split(endpoint.Name, ".")

			// only processing 1 part right now
			if name != strings.ToLower(parts[0]) {
				continue
			}

			split := camelcase.Split(parts[1])

			// find the schema
			schema := spec.Components.Schemas[parts[1]+"Request"]

			if schema == nil {
				fmt.Println(endpoint.Name)
				command.AddCommand(&cobra.Command{
					Use:   strings.ToLower(split[0]),
					Short: "# Unsupported command",
				})
				continue
			}

			var subcommand *cobra.Command

			for i, v := range split {
				cmd := &cobra.Command{
					Use:   strings.ToLower(v),
					Short: "# " + schema.Value.Description,
				}

				// process the args on the last one
				// Lookup process Lookup
				// LookupAddress process Address
				// LookupByPostcode process Postcode
				if i == (len(split) - 1) {
					vals := setFlags(cmd, schema.Value)

					// supply the run command
					cmd.Run = callService(name, parts[1], vals)
				}

				// there is not yet a subcommand for the service
				if subcommand == nil {
					subcommand = cmd
					continue
				}

				commands := subcommand.Commands()

				// there are no subcommands for the endpoint
				if len(commands) == 0 {
					subcommand.AddCommand(cmd)
					continue
				}

				// there are sub commands for the endpoint
				commands[len(commands)-1].AddCommand(cmd)
			}

			command.AddCommand(subcommand)
		}

		commands = append(commands, command)
	}

	return commands
}

func callService(service, endpoint string, vals map[string]interface{}) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		// check for a token
		token := os.Getenv("M3O_API_TOKEN")
		if len(token) == 0 {
			fmt.Fprintln(os.Stderr, "Missing M3O_API_TOKEN environment variable")
			os.Exit(1)
		}

		var req map[string]interface{}

		// marshal the request
		request, err := json.Marshal(vals)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse flags as JSON: %v\n", err)
			os.Exit(1)
		}

		if err := json.Unmarshal([]byte(request), &req); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse request as JSON: %v\n", err)
			os.Exit(1)
		}

		c := client.NewClient(nil)
		c.SetToken(token)

		var rsp json.RawMessage

		if err := c.Call(service, endpoint, req, &rsp); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		var out bytes.Buffer
		defer out.Reset()

		if err := json.Indent(&out, rsp, "", "\t"); err != nil {
			fmt.Fprintln(os.Stderr, "Error while trying to format the response", err)
		}

		fmt.Printf("%v\n", string(out.Bytes()))
	}
}

func setFlags(cmd *cobra.Command, schema *openapi3.Schema) map[string]interface{} {
	vals := map[string]interface{}{}

	for name, property := range schema.Properties {
		switch property.Value.Type {
		case "string":
			val := cmd.Flags().String(name, "", property.Value.Description)
			vals[name] = val
		case "boolean":
			val := cmd.Flags().Bool(name, false, property.Value.Description)
			vals[name] = val
		case "number":
			switch property.Value.Format {
			case "int32":
				val := cmd.Flags().Int32(name, 0, property.Value.Description)
				vals[name] = val
			case "int64":
				val := cmd.Flags().Int64(name, 0, property.Value.Description)
				vals[name] = val
			case "float":
				val := cmd.Flags().Float32(name, 0.0, property.Value.Description)
				vals[name] = val
			case "double":
				val := cmd.Flags().Float64(name, 0.0, property.Value.Description)
				vals[name] = val
			}
		default:
			val := cmd.Flags().String(name, "", property.Value.Description)
			vals[name] = val
		}
	}

	return vals
}

func Execute() {

	// static command list
	exploreSearchCmd.Flags().StringVarP(&Query, "query", "q", "", "The query to search for")
	exploreSearchCmd.MarkFlagRequired("query")
	exploreCmd.AddCommand(exploreListCmd)
	exploreCmd.AddCommand(exploreSearchCmd)

	// client command
	clientCmd.AddCommand(callCmd)
	clientCmd.AddCommand(streamCmd)

	// add the commands
	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(exploreCmd)

	// add the dynamic cli
	commands := generateCLI()
	for _, command := range commands {
		rootCmd.AddCommand(command)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
