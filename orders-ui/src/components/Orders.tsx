import React from "react";
import { useQuery, useSubscription } from "@apollo/client";
import { Order } from "../utils/types";
import { OrderComponent } from "./Order";
import { Grid2 } from "@mui/material";
import { GET_ORDERS, GET_UPDATED_ORDERS } from "../utils/queries";

const Orders: React.FC = () => {
  // const { data: { orders } = {}, loading } = useQuery(GET_ORDERS);
  const { data: { ordersUpdated: orders } = {}, loading } =
    useSubscription(GET_UPDATED_ORDERS);

  if (loading) return <p>Loading...</p>;

  return (
    <>
      <h1>Live Orders</h1>
      <Grid2 container spacing={2}>
        {orders?.map((order: Order) => (
          <OrderComponent order={order} key={order.id} />
        ))}
      </Grid2>
    </>
  );
};

export default Orders;
