from config import config
from src.api import InitController

configuration = config["development"]

controller = InitController()

app = controller.init_app(configuration)

if __name__ == "__main__":
    app.run(host=configuration.APP_HOST, port=configuration.APP_PORT)
