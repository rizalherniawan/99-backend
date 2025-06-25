# Use an official Python runtime as a base image
FROM python:3.11-slim

# Set working directory in container
WORKDIR /app

# Copy requirements and install dependencies
COPY python-libs.txt ./
RUN pip install --no-cache-dir -r python-libs.txt

# Copy the rest of the application code
COPY . .

# Command to run the app (adjust if needed)
CMD ["python", "main.py"]
