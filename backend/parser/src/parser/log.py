import logging

from parser.settings import settings

logging.getLogger().setLevel(settings.log_level)
FORMAT = "[%(asctime)-15s] [%(levelname)s] %(message)s"
logging.basicConfig(format=FORMAT)

logger = logging.getLogger(__name__)
