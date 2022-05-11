from crypt import methods
import json
from flask import Flask, jsonify, request, url_for
from tasks import get_stock_prices
from celery.result import AsyncResult

app = Flask(__name__)


@app.route("/hello", methods=["GET"])
def hello():
    return jsonify({"result": "Hello There!"}), 200


@app.route("/tasks", methods=["POST", "GET"])
def run_scraper():
    response = request.json
    print(response)
    content = json.loads(response)
    print(content)
    stock_list = list(content["symbols"])
    task = get_stock_prices.delay(stock_list)
    return jsonify({"task_status": url_for("get_results", task_id=task.id)}), 202


@app.route("/tasks/<task_id>", methods=["GET"])
def get_results(task_id):
    task_result = AsyncResult(task_id)
    results = {
        "task_id": task_id,
        "task_status": task_result.status,
        "task_result": task_result.result
    }
    return jsonify({"results": results}), 200


if __name__ == "__main__":
    app.run(port=5004)
