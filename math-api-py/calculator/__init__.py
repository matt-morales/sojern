from flask import Flask, jsonify, request


def create_app():
    app = Flask(__name__)

    def get_request_args():
        """Get args from request"""
        data = request.json
        return data.get("numbers"), data.get("q")

    @app.post("/min")
    def post_min():
        numbers, _ = get_request_args()
        _min = min(numbers)
        res = {"min": [_min for n in range(numbers.count(_min))]}
        return jsonify(res), 200

    @app.post("/max")
    def post_max():
        numbers, _ = get_request_args()
        _max = max(numbers)
        res = {"max": [_max for n in range(numbers.count(_max))]}
        return jsonify(res), 200

    @app.post("/avg")
    def post_avg():
        numbers, _ = get_request_args()
        res = {"avg": sum(numbers) / float(len(numbers))}
        return jsonify(res), 200

    @app.post("/median")
    def post_median():
        numbers, _ = get_request_args()
        s_nums = sorted(numbers)
        mid = len(numbers) / 2.0
        res = {
            "median": s_nums[int(mid)]
            if mid % 2
            else (s_nums[int(mid) - 1] + s_nums[int(mid)]) / 2
        }
        return jsonify(res), 200

    @app.post("/percentile")
    def post_percentile():
        numbers, q = get_request_args()
        s_nums = sorted(numbers)
        orank = (q / 100) * len(numbers)
        idx = int(orank) if orank % 1 else int(orank - 1)
        res = {f"{q}_percentile": s_nums[idx]}
        return jsonify(res), 200

    return app
