package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

const (
	LOCATION_RANGE   = 10000.0
	POPULATION_SIZE  = 10000
	DNA_SIZE         = 200
	GENERATION_COUNT = 10000
	MUTATION_CHANCE  = 0.01
	SELECTION_SIZE   = 0.666667
	CROSSOVER_POINT  = 0.5
)

type Gene struct {
	Id int
	X  float64
	Y  float64
}

type Chromosome struct {
	Genes   []Gene
	Fitness float64
}

func initPopulation(populationSize, genesCount int) []Chromosome {
	genes := make([]Gene, genesCount)
	for i := 0; i < genesCount; i++ {
		genes[i] = Gene{
			Id: i,
			X:  rand.Float64() * LOCATION_RANGE,
			Y:  rand.Float64() * LOCATION_RANGE,
		}
	}
	population := make([]Chromosome, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = Chromosome{
			Genes:   genes,
			Fitness: calculateFitness(genes),
		}
	}
	return population
}

func calculateFitness(genes []Gene) float64 {
	totalDistance := 0.0
	for i := 0; i < len(genes)-1; i++ {
		totalDistance += distance(genes[i], genes[i+1])
	}
	totalDistance += distance(genes[len(genes)-1], genes[0])
	return LOCATION_RANGE * LOCATION_RANGE / totalDistance
}

func distance(gene1, gene2 Gene) float64 {
	return math.Sqrt((gene1.X-gene2.X)*(gene1.X-gene2.X) + (gene1.Y-gene2.Y)*(gene1.Y-gene2.Y))
}

func selection(population []Chromosome) []Chromosome {
	sort.Slice(population, func(i, j int) bool {
		return population[i].Fitness > population[j].Fitness
	})
	best := population[:(int)(len(population)*SELECTION_SIZE)]
	return best
}

func crossover(population []Chromosome) []Chromosome {
	for i := 0; i < len(population)/2; i++ {
		parent1 := population[i]
		parent2 := population[len(population)-1-i]
		child1 := make([]Gene, len(parent1.Genes))
		copy(child1, parent1.Genes)
		for j := 0; j < len(parent2.Genes); j++ {
			if !contains(child1, parent2.Genes[j]) {
				child1[len(child1)-1] = parent2.Genes[j]
				break
			}
		}
		population = append(population, Chromosome{
			Genes:   child1,
			Fitness: calculateFitness(child1),
		})
	}

	return population
}

func contains(genes []Gene, gene Gene) bool {
	for _, g := range genes {
		if g.Id == gene.Id {
			return true
		}
	}
	return false
}

func mutation(population []Chromosome) []Chromosome {
	for i := 0; i < len(population); i++ {
		if rand.Float64() <= MUTATION_CHANCE {
			j := rand.Intn(len(population[i].Genes))
			k := rand.Intn(len(population[i].Genes))
			tmp := population[i].Genes[j]
			population[i].Genes[j] = population[i].Genes[k]
			population[i].Genes[k] = tmp
		}
	}
	return population
}

func printPopulation(population []Chromosome) {
	for _, chromosome := range population {
		fmt.Println(chromosome)
	}
}

func main() {
}

