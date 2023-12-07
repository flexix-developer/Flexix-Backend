// healthcheck.js

const express = require('express');
const router = express.Router();

router.get('/health', (req, res) => {
  // Check if the server is running
  const serverStatus = {
    server: 'running',
    timestamp: new Date(),
  };

  // Check if the database connection is healthy
  const dbStatus = {
    database: 'flexixdb',
    status: 'unknown', // You might want to perform an actual health check here
    timestamp: new Date(),
  };

  // You can add more checks as needed

  // Respond with the combined status
  res.status(200).json({
    server: serverStatus,
    database: dbStatus,
  });
});

module.exports = router;
