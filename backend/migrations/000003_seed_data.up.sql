-- Insert user
INSERT INTO users (id, name, email, password, created_at)
VALUES (
  '11111111-1111-1111-1111-111111111111',
  'Test User',
  'test@example.com',
  '$2a$12$3b3y9W8QCZdtBIxn7RaiG.ZtqlVMgUPVtq2d93SmWzfCNV8Izj0ku', -- bcrypt for "password123"
  NOW()
);

-- Insert project
INSERT INTO projects (id, name, description, owner_id, created_at)
VALUES (
  '22222222-2222-2222-2222-222222222222',
  'Sample Project',
  'This is a seeded project',
  '11111111-1111-1111-1111-111111111111',
  NOW()
);

-- Insert tasks
INSERT INTO tasks (id, title, description, status, priority, project_id, assignee_id, created_at, updated_at)
VALUES
(
  '33333333-3333-3333-3333-333333333333',
  'Task 1',
  'First task',
  'todo',
  'high',
  '22222222-2222-2222-2222-222222222222',
  '11111111-1111-1111-1111-111111111111',
  NOW(),
  NOW()
),
(
  '44444444-4444-4444-4444-444444444444',
  'Task 2',
  'Second task',
  'in_progress',
  'medium',
  '22222222-2222-2222-2222-222222222222',
  NULL,
  NOW(),
  NOW()
),
(
  '55555555-5555-5555-5555-555555555555',
  'Task 3',
  'Third task',
  'done',
  'low',
  '22222222-2222-2222-2222-222222222222',
  NULL,
  NOW(),
  NOW()
);