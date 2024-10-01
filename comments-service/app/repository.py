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


class CommentRepository(AbstractRepository):
    def __init__(self, session, model):
        self.session = session
        self.model = model

    def create(self, entity):
        self.session.add(entity)
        self.session.commit()
        return entity

    def list(self, filters=None):
        return self.session.query(self.model).all()

    def get(self, id):
        return self.session.query(self.model).get(id)

    def update(self, id, entity):
        self.session.query(self.model).filter(self.model.id == id).update(entity)
        self.session.commit()
        return entity

    def delete(self, id):
        self.session.query(self.model).filter(self.model.id == id).delete()
