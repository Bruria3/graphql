import { gql } from "@apollo/client";

export const GET_ORDERS = gql`
  query {
    orders {
      id
      name
      status
      createdAt
      quantity
    }
  }
`;

export const GET_UPDATED_ORDERS = gql`
  subscription OrdersUpdated {
    ordersUpdated {
      id
      name
      status
      quantity
      createdAt
    }
  }
`;

export const UPDATE_ORDER = gql`
  mutation UpdateOrder($order: OrderUpdateInput!) {
    updateOrder(order: $order) {
      id
      quantity
    }
  }
`;
