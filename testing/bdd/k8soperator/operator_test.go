package operator_test

import (
	"context"
	"fmt"
	operator "go-training/testing/bdd/k8soperator"
	"testing"

	"github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

type operatorCtx struct{}

func iHaveAClusterWithMinNodesAndMaxNodes(ctx context.Context, min, max int) (context.Context, error) {
	o := operator.NewMemcachedOperator(min, max)
	ctx = context.WithValue(ctx, operatorCtx{}, o)
	return ctx, nil
}

func iScaleDownTheClusterByNodes(ctx context.Context, arg1 int) (context.Context, error) {
	o := ctx.Value(operatorCtx{}).(*operator.MemcachedOperator)
	o.Scale(-arg1)
	ctx = context.WithValue(ctx, operatorCtx{}, o)

	return ctx, nil
}

func iScaleUpTheClusterByNodes(ctx context.Context, arg1 int) (context.Context, error) {
	o := ctx.Value(operatorCtx{}).(*operator.MemcachedOperator)
	o.Scale(arg1)
	ctx = context.WithValue(ctx, operatorCtx{}, o)

	return ctx, nil
}

func theClusterShouldHaveNodes(ctx context.Context, expected int) error {
	o := ctx.Value(operatorCtx{}).(*operator.MemcachedOperator)

	if o.CountNodes() != expected {
		return fmt.Errorf("expected %d nodes, got %d", expected, o.CountNodes())
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I have a cluster with min nodes (\d+) and max nodes (\d+)$`, iHaveAClusterWithMinNodesAndMaxNodes)
	ctx.Step(`^I scale down the cluster by (\d+) nodes$`, iScaleDownTheClusterByNodes)
	ctx.Step(`^I scale up the cluster by (\d+) nodes$`, iScaleUpTheClusterByNodes)
	ctx.Step(`^the cluster should have (\d+) nodes$`, theClusterShouldHaveNodes)
}
