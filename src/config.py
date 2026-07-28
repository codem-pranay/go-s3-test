import os
import sys

class Config:
    """Application configuration management"""
    
    def __init__(self):
        self.database_url = os.environ.get('DATABASE_URL')
        
    def validate(self):
        """Validate required environment variables"""
        if not self.database_url:
            raise ValueError(
                "DATABASE_URL environment variable is not set. "
                "Please configure it in CodeBuild environment variables or buildspec.yml"
            )
        return True
    
    def get_database_url(self):
        """Get database connection URL"""
        return self.database_url


def get_config():
    """Factory function to get validated configuration"""
    config = Config()
    config.validate()
    return config

