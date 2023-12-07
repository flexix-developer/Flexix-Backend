const express = require('express');
const bodyParser = require('body-parser');
const mysql = require('mysql');
const healthCheckRouter = require('../healthcheck');

const app = express();
const port = 5000;

// Middleware
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

// MySQL Connection
const db = mysql.createConnection({
  host: 'db',
  user: 'root',
  password: 'root',
  database: 'flexixdb'
});

db.connect((err) => {
  if (err) {
    console.error('Database connection failed: ' + err.stack);
    process.exit(1); // Terminate the application
  }
  console.log('Connected to database');
});


// CRUD Operations

// Create User
app.post('/users', (req, res) => {
  const { Fname, Lname, Email, Pass } = req.body;

  const createUserSql = 'INSERT INTO users (Fname, Lname, Email, Pass) VALUES (?, ?, ?, ?)';
  db.query(createUserSql, [Fname, Lname, Email, Pass], (err, userResult) => {
    if (err) {
      console.error('Error creating user:', err);
      return res.status(500).send('Error creating user');
    }

    const userId = userResult.insertId;
    const createWorkspaceSql = 'INSERT INTO workspace (Users_ID) VALUES (?)';
    db.query(createWorkspaceSql, [userId], (err, workspaceResult) => {
      if (err) {
        console.error('Error creating workspace:', err);
        return res.status(500).send('Error creating workspace');
      }

      res.status(201).send('User created successfully with associated Workspace');
    });
  });
});

// Read Users
app.get('/users', (req, res) => {
  const sql = 'SELECT * FROM users';
  db.query(sql, (err, results) => {
    if (err) {
      console.error('SQL error:', err);
      return res.status(500).send('Error fetching users');
    }
    res.status(200).json(results);
  });
});

// Update User
app.put('/users/:id', (req, res) => {
  const userId = req.params.id;
  const { Fname, Lname, Email, Pass } = req.body;
  const sql = 'UPDATE users SET Fname=?, Lname=?, Email=?, Pass=? WHERE ID=?';
  db.query(sql, [Fname, Lname, Email, Pass, userId], (err, result) => {
    if (err) {
      console.error('Error updating user:', err);
      return res.status(500).send('Error updating user');
    }

    res.status(200).send('User updated successfully');
  });
});

// Delete User
app.delete('/users/:id', (req, res) => {
  const userId = req.params.id;
  const sql = 'DELETE FROM users WHERE ID=?';
  db.query(sql, [userId], (err, result) => {
    if (err) {
      console.error('Error deleting user:', err);
      return res.status(500).send('Error deleting user');
    }

    res.status(200).send('User deleted successfully');
  });
});

// Health Check Route
app.use('/', healthCheckRouter);

// Start server
app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
