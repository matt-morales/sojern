import json


def test_post_min(client):
    numbers = [10, 20, 15, -5, 0, 20]
    request_data = {
        "numbers": numbers,
        "q": len(numbers),
    }
    resp = client.post("/min", json=request_data)
    res = resp.json
    assert resp.status_code == 200
    assert "min" in res
    assert res["min"] == [-5]


def test_post_max(client):
    numbers = [10, 20, 15, -5, 0, 20]
    request_data = {
        "numbers": numbers,
        "q": len(numbers),
    }
    resp = client.post("/max", json=request_data)
    res = resp.json
    assert resp.status_code == 200
    assert "max" in res
    assert res["max"] == [20, 20]


def test_post_avg(client):
    numbers = [10, 20, 15, -5, 0, 20]
    request_data = {
        "numbers": numbers,
        "q": len(numbers),
    }
    resp = client.post("/avg", json=request_data)
    res = resp.json
    assert resp.status_code == 200
    assert "avg" in res
    assert res["avg"] == 10


def test_post_median(client):
    numbers = [10, 20, 15, -5, 0, 20]
    request_data = {
        "numbers": numbers,
        "q": len(numbers),
    }
    resp = client.post("/median", json=request_data)
    res = resp.json
    assert resp.status_code == 200
    assert "median" in res
    assert res["median"] == 15


def test_post_percentile(client):
    numbers = [10, 20, 15, -5, 0, 20]
    q = 61
    request_data = {
        "numbers": numbers,
        "q": q,
    }
    resp = client.post("/percentile", json=request_data)
    res = resp.json
    assert resp.status_code == 200
    assert f"{q}_percentile" in res
    assert res[f"{q}_percentile"] == 15
