import json
import random
import web


class GeneratorServer(object):
    """
    This class is responsible for initializing the urls and web server.
    """
    def __init__(self, port=80, host='0.0.0.0'):
        inst = Generator()
        urls = inst.urls()
        app = web.application(urls, globals())
        web.httpserver.runsimple(app.wsgifunc(), (host, port))


class GenerateIt:
    @staticmethod
    def GET():
        a = {}
        i = 1
        while i < 10:
            a['field'+str(i)] = random.randrange(1, 500*i)
            i += 1
        return json.dumps(a)


class Generator:
    """
    This class is for defining things needed to start up.
    """

    @staticmethod
    def urls():
        urls = (
            '/genit', GenerateIt
        )
        return urls

if __name__ == '__main__':
    GeneratorServer().app.run()
