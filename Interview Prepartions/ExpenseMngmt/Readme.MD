To visualize the code flow and table structures for the expense-sharing application, let's break it down into a diagram and a corresponding explanation.

### Visual Code Flow Diagram

The flowchart below outlines the key steps and interactions in the application:

```
+-----------------------+
|       Start           |
+-----------------------+
          |
          v
+-----------------------+
|   Initialize Expense  |
|     Manager           |
+-----------------------+
          |
          v
+-----------------------+     +------------------+
|   Add User            | --> |  Generate Unique  |
|                       |     |  User ID (UUID)   |
+-----------------------+     +------------------+
          |                             |
          v                             v
+-------------------------------+    +------------------+
| Store User in Users Map       |    | Return User ID    |
+-------------------------------+    +------------------+
          |
          v
+-------------------------------+
|   Add Expense                 |
|                               |
|  1. Select Split Strategy     |
|  2. Calculate Share           |
+-------------------------------+
          |
          v
+-------------------------------+
| Update Balances in Balances   |
| Map                           |
+-------------------------------+
          |
          v
+-------------------------------+
| Show Balances                 |
|  1. For All Users             |
|  2. For Specific User         |
+-------------------------------+
          |
          v
+-------------------------------+
|          End                  |
+-------------------------------+
```

### Explanation:
1. **Start**: The App initializing `ExpenseManager`.

2. **Add User**:
   - Users are added by generating a unique user ID using UUID.
   - The user details (ID, name, email, phone) are stored in the `users` map.

3. **Add Expense**:
   - An expense is added by selecting the appropriate splitting strategy (`EqualSplit`, `ExactSplit`, or `PercentageSplit`).
   - The selected strategy calculates the share for each user involved.
   - Balances between users are updated based on the calculated shares.

4. **Show Balances**:
   - Balances can be displayed for all users or a specific user.
   - The application shows who owes how much to whom.


### Table Structures

In the context of this application, two primary data structures (`users` and `balances`) are used.

#### 1. **Users Table (Map)**
| Field      | Type   | Description                          |
|------------|--------|--------------------------------------|
| `ID`       | String | Unique identifier for the user (UUID)|
| `Name`     | String | Name of the user                     |
| `Email`    | String | Email address of the user            |
| `Phone`    | String | Phone number of the user             |

- **Storage**: Stored in the `users` map (`map[string]*User`).
- **Purpose**: To store user information and uniquely identify each user.

#### 2. **Balances Table (Nested Map)**
| Field 1 (`Owed By`) | Field 2 (`Owed To`) | Amount (Float64) | Description                                     |
|---------------------|---------------------|------------------|-------------------------------------------------|
| `UserID1`           | `UserID2`           | `Amount`         | Amount that `UserID1` owes to `UserID2`         |
| `UserID2`           | `UserID1`           | `Amount`         | Amount that `UserID2` owes to `UserID1`         |

- **Storage**: Stored in the `balances` map (`map[string]map[string]float64`).
- **Purpose**: To keep track of the balances between users, indicating who owes how much to whom.

### Example Tables with Data

Example transactions:

#### Users Table

| ID                                 | Name  | Email               | Phone       |
|------------------------------------|-------|---------------------|-------------|
| `1f3d2e1c-8c7b-4d13-b23b-abcdef123`| User1 | user1@example.com   | 1234567890  |
| `4d2b3a4c-8a9d-41b2-b23b-abcd1234ef`| User2 | user2@example.com   | 0987654321  |
| `5e3f6b1c-7f3e-4e0f-b89b-abcf123efg`| User3 | user3@example.com   | 1122334455  |

#### Balances Table

| Owed By                            | Owed To                             | Amount  |
|------------------------------------|-------------------------------------|---------|
| `4d2b3a4c-8a9d-41b2-b23b-abcd1234ef`| `1f3d2e1c-8c7b-4d13-b23b-abcdef123` | 620.00  |
| `5e3f6b1c-7f3e-4e0f-b89b-abcf123efg`| `1f3d2e1c-8c7b-4d13-b23b-abcdef123` | 1130.00 |
| `1f3d2e1c-8c7b-4d13-b23b-abcdef123`| `4d2b3a4c-8a9d-41b2-b23b-abcd1234ef`| 230.00  |

