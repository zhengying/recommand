#recommand
=========
A golang library for calculating the similarity.

you can choose follow algorithm
>* cosine similarity
>* precision similarity
>* euclidean distance
 
 
 ###How to use
 
 ``` go
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
	for _, item := range sim.SimilarityList(similarity.Type_Cosine_similarity, "P1", 7) {
		log.Println(item.NodeFieldScores)
	}
	```
