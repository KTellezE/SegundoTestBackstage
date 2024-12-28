from src.persistence.Db.database import db

class FunctionsUtil():
    
    def DeleteAllRecords():
        meta = db.metadata
        for table in reversed(meta.sorted_tables):
            db.session.execute(table.delete())
        db.session.commit()
    
