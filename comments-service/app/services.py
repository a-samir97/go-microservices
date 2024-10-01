from repository import CommentRepository


class CommentService:
    def __init__(self, comment_repo: CommentRepository):
        self.comment_repo = comment_repo

    def get_comments(self):
        return self.comment_repo.list()

    def get_comment(self, comment_id):
        return self.comment_repo.get(comment_id)

    def add_comment(self, comment):
        return self.comment_repo.create(comment)

    def delete_comment(self, comment_id):
        return self.comment_repo.delete(comment_id)

    def update_comment(self, comment_id, comment):
        return self.comment_repo.update(comment_id, comment)