package main

import (
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"

	"github.com/heimathafen/hafen/proto/recipe"
	"github.com/heimathafen/hafen/rpc/grpc/search"
	"google.golang.org/grpc"
)

type Service struct{}

func (*Service) Call(ctx context.Context, req *search.SearchRequest) (*search.Results, error) {
	if !strings.Contains("docker-compose", req.Query) {
		return &search.Results{}, nil
	}

	return &search.Results{
		Recipes: []*recipe.Recipe{
			{
				Name:             "docker-compose",
				Version:          "0.1.0",
				RecipeGitUrl:     "https://github.com/heimathafen/library",
				RecipeConfigPath: "/docker-compose/config.json",
				RecipeBuildPath:  "/docker-compose/build/",
				RecipeRunPath:    "/docker-compose/run/",
				RecipeHttpUrl:    "https://github.com/heimathafen/library/tree/master/docker-compose/",
				GitUrl:           "https://github.com/docker/compose",
			},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":30100")
	if err != nil {
		log.Fatalf("Failed to listen on tcp/:30100; Error: %s", err)
	}
	log.Printf("Stub Search Service is listening on *:30100")

	s := grpc.NewServer()
	search.RegisterSearchServer(s, &Service{})
	s.Serve(lis)
}
