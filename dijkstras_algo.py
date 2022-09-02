from heapq import *

# Time: O((v+e)*log(v)) | Space: O(v)
def dikstrasAlgorithm(start, edges):
    numberOfVertices = len(edges)
    distances = [float("inf") for _ in range( numberOfVertices)]
    distances[start] = 0
    
    minDistancesHeap = [(0, start)] # (distance, vertexIdx)

    # visited = set()

    while minDistancesHeap:
        currentMinDistance, vertexIdx = heappop(minDistancesHeap)

        # if vertexIdx in visited:
        #     continue

        # visited.add(vertexIdx)

        for edge in edges[vertexIdx]:
            adjacentVertexIdx, distanceToAdjacentVertex = edge[0], edge[1]

            # if adjacentVertexIdx in visited:
            #     continue

            newPathDistance = currentMinDistance + distanceToAdjacentVertex
            currentAdjacentVertexDistance = distances[adjacentVertexIdx]
            if newPathDistance < currentAdjacentVertexDistance:
                distances[adjacentVertexIdx] = newPathDistance
                heappush(minDistancesHeap, (newPathDistance, adjacentVertexIdx))

    return distances

if __name__ == "__main__":
    edges = [
        # for vertex 0
        [
            [1, 2], [3, 6]
        ],
         # for vertex 1
        [
            [0, 2], [2, 5]
        ],
         # for vertex 2
        [
            [1, 5], [3, 7], [4, 6], [5, 9]
        ],
         # for vertex 3
        [
            [0, 6], [2, 7], [4, 10]
        ],
         # for vertex 4
        [
            [2, 6], [3, 10], [5, 6]
        ],
         # for vertex 5
        [
            [2, 9], [4, 6]
        ],
    ]
    # source/start node is 0
    print("distances", dikstrasAlgorithm(0, edges))

    # output: 'distances', [0, 2, 7, 6, 13, 16])
