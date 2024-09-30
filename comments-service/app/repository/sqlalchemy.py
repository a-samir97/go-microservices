from .db import AbstractRepository


class CommentRepository(AbstractRepository):
    def __init__(self, session, model):
        self.session = session
        self.model = model

    def create(self, entity):
        self.session.add(entity)
        return entity

    def list(self, filters=None):
        return self.session.query(self.model).all()

    def get(self, id):
        return self.session.query(self.model).get(id).first()

    def update(self, id, entity):
        self.session.query(self.model).filter(self.model.id == id).update(entity)
        return entity

    def delete(self, id):
        self.session.query(self.model).filter(self.model.id == id).delete()
