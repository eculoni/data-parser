# utils.py

import logging
import os
import json
import re

from dataclasses import dataclass
from typing import Dict, List, Optional

# Set up logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

@dataclass
class Config:
    input_dir: str
    output_dir: str
    log_level: int

def load_config(config_path: str) -> Config:
    with open(config_path, 'r') as f:
        config_data = json.load(f)
    return Config(**config_data)

def validate_config(config: Config) -> bool:
    if not os.path.exists(config.input_dir):
        logger.error(f"Input directory '{config.input_dir}' does not exist")
        return False
    if not os.path.exists(config.output_dir):
        logger.error(f"Output directory '{config.output_dir}' does not exist")
        return False
    return True

def parse_json_file(file_path: str) -> Dict:
    with open(file_path, 'r') as f:
        data = json.load(f)
    return data

def extract_data(data: Dict) -> List[Dict]:
    # Simple example, replace with actual data extraction logic
    return [{**d, 'extracted_field': d['field']} for d in data['data']]

def validate_data(data: List[Dict]) -> bool:
    for item in data:
        if not isinstance(item, dict):
            logger.error(f"Invalid data format: {item}")
            return False
        if 'field' not in item:
            logger.error(f"Missing required field: {item}")
            return False
    return True