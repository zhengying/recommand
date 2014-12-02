package main

import (
	"github.com/zhengying/recommand/similarity"
	"log"
)

func main() {

	sim := similarity.NewSimilarity()

	node := similarity.Node{
		NodeID:          "P1",
		NodeFieldScores: []float64{180.0, 75.0, 40.0},
	}
	sim.AddNode(node)
	node = similarity.Node{
		NodeID:          "P2",
		NodeFieldScores: []float64{140.0, 20, 0.0},
	}
	sim.AddNode(node)
	node = similarity.Node{
		NodeID:          "P3",
		NodeFieldScores: []float64{130.0, 75.0, 20.0},
	}
	sim.AddNode(node)
	node = similarity.Node{
		NodeID:          "P4",
		NodeFieldScores: []float64{110.0, 75.0, 20.0},
	}
	sim.AddNode(node)
	node = similarity.Node{
		NodeID:          "P5",
		NodeFieldScores: []float64{170.0, 15.0, 20.0},
	}
	sim.AddNode(node)
	node = similarity.Node{
		NodeID:          "P6",
		NodeFieldScores: []float64{150.0, 35.0, 20.0},
	}
	sim.AddNode(node)
	node = similarity.Node{
		NodeID:          "P7",
		NodeFieldScores: []float64{80.0, 5.0, 6.0},
	}
	sim.AddNode(node)
	node = similarity.Node{
		NodeID:          "P8",
		NodeFieldScores: []float64{70.0, 25.0, 8.0},
	}
	sim.AddNode(node)

	log.Println("dest: {180.0, 75.0, 40.0}")

	log.Println("--cosine similarity--")
	for _, item := range sim.SimilarityList(similarity.Type_Cosine_similarity, "P8", 7) {
		log.Println(item.NodeFieldScores, item.Score)
	}

	log.Println("--distance similarity--")
	for _, item := range sim.SimilarityList(similarity.Type_Distance_similarity, "P8", 7) {
		log.Println(item.NodeFieldScores, item.Score)
	}

	log.Println("--precision similarity--")
	for _, item := range sim.SimilarityList(similarity.Type_precision_similarity, "P8", 7) {
		log.Println(item.NodeFieldScores, item.Score)
	}
}
