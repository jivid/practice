def path_exists_bfs(graph, start, end):
    all_vertices = graph.vertices
    for _, vert in all_vertices.items():
        vert.state = States.UNSEEN

    start = graph.find_vertex(start)
    if start is None:
        return False

    if start.value == end:
        return True

    found = False
    vertices = [start]

    all_done = lambda x: all([c.state == States.DONE for c in x])
    while not all_done(vertices):
        for _, v in enumerate(vertices):
            if v.state == States.DONE:
                continue

            for c in v.edges:
                if c == end:
                    found = True
                    break

                vert = graph.find_vertex(c)
                if vert.state == States.UNSEEN:
                    vert.state = States.SEEN
                    vertices.append(vert)

            v.state = States.DONE

    return found
