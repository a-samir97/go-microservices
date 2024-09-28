from .db import AbstractRepository


class InfluxDBRepository(AbstractRepository):
    def create(self, entity):
        return super().create(entity)

    def list(self, filters=None):
        return super().list(filters)

    def get(self, id):
        return super().get(id)

    def update(self, entity):
        return super().update(entity)

    def delete(self, id):
        return super().delete(id)
