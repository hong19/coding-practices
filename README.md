# coding-practices

## Setup 

```bash
python -m venv .venv
source .venv/bin/activate

pip install -e .\[dev\] 
```

### Test 
```bash
source .venv/bin/activate

pytest leetcode/9. Palindrome Number/test.py
# or 
cd leetcode/9. Palindrome Number/
pytest test.py
```
