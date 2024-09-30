from abc import ABC, abstractmethod


class AbstractRepository(ABC):
    @abstractmethod
    def get(self, id):
        raise NotImplementedError

    @abstractmethod
    def list(self, filters=None):
        raise NotImplementedError

    @abstractmethod
    def create(self, entity):
        raise NotImplementedError

    @abstractmethod
    def update(self, entity):
        raise NotImplementedError

    @abstractmethod
    def delete(self, id):
        raise NotImplementedError
