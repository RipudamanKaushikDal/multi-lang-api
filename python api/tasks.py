from celery import Celery, Task
from stock_scraper import YahooFinanceScraper
import asyncio
import os


celery = Celery(__name__)
celery.conf.broker_url = os.environ.get(
    "CELERY_BROKER_URL", "redis://localhost:6379")
celery.conf.result_backend = os.environ.get(
    "CELERY_RESULT_BACKEND", "redis://localhost:6379")


@celery.task(name="get_stock_prices")
def get_stock_prices(stocklist: list) -> Task:
    scraper = YahooFinanceScraper(symbol_list=stocklist)
    result_list = asyncio.run(scraper.fetch_prices())
    print(result_list)
    return result_list
