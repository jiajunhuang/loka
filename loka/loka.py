import os


class Meta(type):
    def __new__(cls, name, bases, namespace, **kwds):
        # set config
        for k, v in namespace.items():
            if k.startswith("_"):
                continue
            envv = os.getenv(k, v)
            namespace[k] = type(v)(envv)

        result = type.__new__(cls, name, bases, dict(namespace))

        return result


class LokaConfig(metaclass=Meta):
    pass
