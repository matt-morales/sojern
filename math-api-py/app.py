from flask import (
    Flask,
    jsonify,
    request
)

app = Flask(__name__)


def get_request_parameters():
    data = request.json
    return data.get("numbers"), data.get("q")

@app.post("/min")
def post_min():
    numbers, _ = get_request_parameters()
    _min = min(numbers)
    res = {"min": [_min for n in range(numbers.count(_min))]}
    return jsonify(res), 200

@app.post("/max")
def post_max():
    numbers, _ = get_request_parameters()
    _max = max(numbers)
    res = {"min": [_max for n in range(numbers.count(_max))]}
    return jsonify(res), 200

@app.post("/avg")
def post_avg():
    numbers, _ = get_request_parameters()
    res = {"avg": sum(numbers) / float(len(numbers))}
    res = {}
    return jsonify(res), 200

@app.post("/median")
def post_median():
    numbers, _ = get_request_parameters()
    s_nums = sorted(numbers)
    mid = len(numbers) / 2.0
    res = {"median": s_nums[int(mid)] if mid % 2 else (s_nums[int(mid) - 1] + s_nums[int(mid)]) / 2}
    return jsonify(res), 200

@app.post("/percentile")
def post_percentile():
    numbers, q = get_request_parameters()
    s_nums = sorted(numbers)
    orank = (q / 100) * len(numbers)
    idx = int(orank) if orank % 1 else int(orank - 1)
    res = { "orank": orank, "percentile": s_nums[idx] }
    return jsonify(res), 200
