# API Reference

## Payments

| Name                              | URL                                  | Method |
| --------------------------------- | ------------------------------------ | ------ |
| Get payments                      | /api/payments                        | GET    |
| Get payments submitted this month | /api/payments/:year/:month           | GET    |
| Get payments for a category       | /api/payments/:year/:month/:category | GET    |
| Submit a payment                  | /api/payments/:year/:month/:category | POST   |
| Amend a payment                   | /api/payment/:id                     | UPDATE |
| Remove a payment                  | /api/payment/:id                     | DELETE |

## Automatic payments

| Name                     | URL                | Method |
| ------------------------ | ------------------ | ------ |
| Get automatic payments   | /api/automatic     | GET    |
| Add automatic payment    | /api/automatic     | POST   |
| Amend automatic payment  | /api/automatic/:id | UPDATE |
| Remove automatic payment | /api/automatic/:id | DELETE |

## Categories

| Name                      | URL                          | Method |
| ------------------------- | ---------------------------- | ------ |
| Get categories            | /api/categories              | GET    |
| Add a category            | /api/categories              | POST   |
| Amend a category          | /api/category/:id            | UPDATE |
| Get categories this month | /api/categories/:year/:month | GET    |

## Income

| Name         | URL         | Method |
| ------------ | ----------- | ------ |
| Get income   | /api/income | GET    |
| Amend income | /api/income | UPDATE |

## Summary

| Name    | URL                       | Method |
| ------- | ------------------------- | ------ |
| Summary | /api/summary/:year/:month | GET    |
