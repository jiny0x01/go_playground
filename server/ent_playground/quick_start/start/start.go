package main

import (
	"context"
	"ent_playground/ent"
	"ent_playground/ent/user"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/examples/start/ent/car"
	"golang.org/x/net/context"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("jiny").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("jiny")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	jiny, err := client.User.
		Create().
		SetAge(30).
		SetName("jiny").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", jiny)
	return jiny, nil
}

func QueryCars(ctx context.Context, jiny *ent.User) error {
	cars, err := jiny.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println("returned cars:", cars)

	ford, err := jiny.QueryCars().
		Where(car.Model("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println("ford")
	return nil
}

func QueryCarUsers(ctx context.Context, jiny *ent.User) error {
	cars, err := jiny.QueryCars.All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}

	for _, ca := range cars {
		owner, err := ca.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %w", ca.Model, err)
		}
		log.Printf("car %q owner: %q\n", ca.Model, owner.Name)
	}
	return nil
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	   a8m, err := client.User.
        Create().
        SetAge(30).
        SetName("Ariel").
        Save(ctx)
    if err != nil {
        return err
    }
    neta, err := client.User.
        Create().
        SetAge(28).
        SetName("Neta").
        Save(ctx)
    if err != nil {
        return err
    }
    // Then, create the cars, and attach them to the users in the creation.
    err = client.Car.
        Create().
        SetModel("Tesla").
        SetRegisteredAt(time.Now()). // ignore the time in the graph.
        SetOwner(a8m).               // attach this graph to Ariel.
        Exec(ctx)
    if err != nil {
        return err
    }
    err = client.Car.
        Create().
        SetModel("Mazda").
        SetRegisteredAt(time.Now()). // ignore the time in the graph.
        SetOwner(a8m).               // attach this graph to Ariel.
        Exec(ctx)
    if err != nil {
        return err
    }
    err = client.Car.
        Create().
        SetModel("Ford").
        SetRegisteredAt(time.Now()). // ignore the time in the graph.
        SetOwner(neta).              // attach this graph to Neta.
        Exec(ctx)
    if err != nil {
        return err
    }
    // Create the groups, and add their users in the creation.
    err = client.Group.
        Create().
        SetName("GitLab").
        AddUsers(neta, a8m).
        Exec(ctx)
    if err != nil {
        return err
    }
    err = client.Group.
        Create().
        SetName("GitHub").
        AddUsers(a8m).
        Exec(ctx)
    if err != nil {
        return err
    }
    log.Println("The graph was created successfully")
    return nil
}

func QueryGithub(ctx context.Context, client *ent. Client) error {
	cars, err := client.Group.
		Qeury().
		Where(group.Name("Github")).
		QueryUsers().
		QeuryCars().
		All(ctx)
	if err != nil {
		return fmt.Errorf("failedd getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	return nil
}


func QueryArielCars(ctx context.Context, client *ent. Client) error {
	jiny := client.User.
		Qeury().
		Where(
			user.HasCars(),
			user.Name("Ariel"),
		).
		OnlyX(ctx)
	cars, err := jiny.
		QueryGroup().
		QeuryUsers().
		QueryCars().
		Where(
			car.Not(
				car.Model("Mazda"),
			),
		).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	return nil
}

func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups, err := client.Group.
		Query().
		Where(group.HasUsers()).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting group: %w", err)
	}
	log.Println("groups returned:", groups")
	return nil
}
