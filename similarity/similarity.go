/*

package main

import (
	"log"
	"zhengying/similarity"
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
	for _, item := range sim.SimilarityList(similarity.Type_Cosine_similarity, "P1", 7) {
		log.Println(item.NodeFieldScores)
	}
}

*/

package similarity

import (
	"log"
	"math"
	//"sort"
)

type SimilarityAlgorithmType int

const (
	Type_Cosine_similarity    = 0
	Type_Distance_similarity  = 1
	Type_precision_similarity = 2
)

// input Node
type Node struct {
	NodeID          string
	NodeFieldScores []float64
}

type NodeOutput struct {
	Node
	Score float64
}

type Similarity struct {
	nodes []Node
}

func NewSimilarity() *Similarity {
	return &Similarity{
		make([]Node, 0, 10),
	}
}

// store index and score
type scoreObject struct {
	index      int
	scoreValue float64
}

// score sorter implement sort interface
type scoreSorter []scoreObject

func (c scoreSorter) Len() int {
	return len(c)
}

func (c scoreSorter) Less(i, j int) bool {
	return c[i].scoreValue < c[j].scoreValue
}

func (c scoreSorter) Swap(i, j int) {
	c[i].index, c[j].index = c[j].index, c[i].index
	c[i].scoreValue, c[j].scoreValue = c[j].scoreValue, c[i].scoreValue
}

func (self *Similarity) AddNode(node Node) {
	self.nodes = append(self.nodes, node)
}

// external function for caller
func (self Similarity) SimilarityList(simType SimilarityAlgorithmType, nodeID string, returnCount int) []NodeOutput {

	if returnCount >= len(self.nodes) {
		panic("list count must less then nodes count-1")
	}

	scores := make([]scoreObject, len(self.nodes))

	srcIndex := -1
	for index, node := range self.nodes {
		if node.NodeID == nodeID {
			srcIndex = index
			break
		}
	}

	if srcIndex == -1 {
		return nil
	}

	for i := 0; i < len(self.nodes); i++ {

		theScore := 0.0

		switch simType {
		case Type_Cosine_similarity:
			theScore = cosine(self.nodes[srcIndex].NodeFieldScores, self.nodes[i].NodeFieldScores)
		case Type_Distance_similarity:
			theScore = distance(self.nodes[srcIndex].NodeFieldScores, self.nodes[i].NodeFieldScores)
		case Type_precision_similarity:
			theScore = precision(self.nodes[srcIndex].NodeFieldScores, self.nodes[i].NodeFieldScores)
		default:
			panic("error Similarity Algorithm Type")
		}

		scores[i].scoreValue = theScore
		scores[i].index = int(i)
	}

	scores = nearest(srcIndex, returnCount, scores)

	retNodes := make([]NodeOutput, 0, 10)
	for i := 0; i < len(scores); i++ {

		if scores[i].index == srcIndex {
			continue
		}

		retNodes = append(retNodes, NodeOutput{self.nodes[scores[i].index], scores[i].scoreValue})
	}

	return retNodes
}

func avg(A []float64, isCalculateNoneValue bool) (calcCount int, avgValue float64) {
	nAvgCount := 0
	sum := 0.0

	for i := 0; i < len(A); i++ {
		if !isCalculateNoneValue && A[i] == 0 {
			continue
		}
		nAvgCount++
		sum += A[i]
	}

	return nAvgCount, sum / float64(nAvgCount)
}

func precision(A []float64, B []float64) float64 {
	lenA := len(A)
	lenB := len(B)

	if lenA != lenB || lenA == 0 || lenB == 0 {
		panic("two vectors must same point")
	}

	_, avgA := avg(A, false)
	_, avgB := avg(B, false)
	sumAvgProduct := 0.0
	sumSqrtDiffA := 0.0
	sumSqrtDiffB := 0.0
	for i := 0; i < lenA; i++ {
		sumAvgProduct += (A[i] - avgA) * (B[i] - avgB)
		sumSqrtDiffA += (A[i] - avgA) * (A[i] - avgA)
		sumSqrtDiffB += (B[i] - avgB) * (B[i] - avgB)
	}
	return sumAvgProduct / (math.Sqrt(sumSqrtDiffA) * math.Sqrt(sumSqrtDiffB))
}

// cosines similarity algorithm
func cosine(A []float64, B []float64) float64 {
	lenA := len(A)
	lenB := len(B)

	if lenA != lenB || lenA == 0 || lenB == 0 {
		panic("two vectors must same point")
	}

	dotProduct := 0.0
	squareSumA := 0.0
	squareSumB := 0.0

	for i := 0; i < lenA; i++ {
		dotProduct += A[i] * B[i]
		squareSumA += A[i] * A[i]
		squareSumB += B[i] * B[i]
	}

	return dotProduct / (math.Sqrt(squareSumA) * math.Sqrt(squareSumB))
}

// euclidean distance similarity algorithm
func distance(A []float64, B []float64) float64 {

	log.Println("distance.....")

	lenA := len(A)
	lenB := len(B)

	if lenA != lenB || lenA == 0 || lenB == 0 {
		log.Println("Error format")
		return 0.0
	}

	disSquareSum := 0.0
	for i := 0; i < lenA; i++ {
		disSquareSum += (A[i] - B[i]) * (A[i] - B[i])
	}

	return math.Sqrt(disSquareSum)
}

// a quick sort for scoreObject, the compare value of array is abs(scoreObject.scoreValue)
func qabsSort(arr []scoreObject, start int, end int, originNode scoreObject) {
	var (
		key  scoreObject = arr[start]
		low  int         = start
		high int         = end
	)

	realNode := func(n scoreObject) float64 {
		return math.Abs(n.scoreValue - originNode.scoreValue)
	}

	for {
		for low < high {
			if realNode(arr[high]) < realNode(key) {
				arr[low] = arr[high]
				break
			}
			high--
		}
		for low < high {
			if realNode(arr[low]) > realNode(key) {
				arr[high] = arr[low]
				break
			}
			low++
		}
		if low >= high {
			arr[low] = key
			break
		}
	}
	if low-1 > start {
		qabsSort(arr, start, low-1, originNode)
	}
	if high+1 < end {
		qabsSort(arr, high+1, end, originNode)
	}
}

// get the ordered nearest list from a sort list
func nearest(destIndex int, retCount int, nodes []scoreObject) []scoreObject {
	qabsSort(nodes, 0, len(nodes)-1, nodes[destIndex])
	return nodes[:retCount+1]
}
