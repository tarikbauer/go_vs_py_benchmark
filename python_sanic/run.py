import os
import asyncio
from sanic import Sanic
from sanic.request import Request
from sanic.exceptions import NotFound
from sanic.response import HTTPResponse, text

app = Sanic(__name__)

async def sleep(seconds: int) -> int:
    await asyncio.sleep(seconds)
    return seconds

@app.exception(NotFound)
async def handling_not_found(request: Request, exception: Exception):
    return text("Not Found!", 404)

@app.route("/api", methods=["GET"])
async def handler(request: Request) -> HTTPResponse:
    tasks = []
    for params in request.query_args:
        if params[0] == "t":
            for value in params[1].split(","):
                try:
                    seconds = int(value)
                except ValueError:
                    return text("Invalid Parameters!", 400)
                tasks.append(asyncio.create_task(sleep(seconds)))
            value = 0
            for task in tasks:
                value += await task
            return text(value)
    return text("Not Found!", 404)


if __name__ == "__main__":
    app.run(os.environ["SANIC_HOST"], int(os.environ["SANIC_PORT"]),workers=int(os.environ["SANIC_WORKERS"]))
