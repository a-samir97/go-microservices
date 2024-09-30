class CommnetService:
    def __init__(self, comment_repo):
        self.comment_repo = comment_repo

    def get_comments(self, user_id):
        return self.comment_repo.get_comments(user_id)

    def add_comment(self, user_id, comment):
        return self.comment_repo.add_comment(user_id, comment)

    def delete_comment(self, user_id, comment_id):
        return self.comment_repo.delete_comment(user_id, comment_id)
