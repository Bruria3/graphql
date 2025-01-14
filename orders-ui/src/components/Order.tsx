import {
  Button,
  Card,
  CardActions,
  CardContent,
  Grid2,
  TextField,
  Typography,
} from "@mui/material";
import { useState } from "react";
import { Order } from "../utils/types";
import { useMutation } from "@apollo/client";
import { UPDATE_ORDER } from "../utils/queries";

interface OrderProps {
  order: Order;
}

export const OrderComponent = ({ order }: OrderProps) => {
  const [quantity, setQuantity] = useState(order.quantity);

  const [updateOrderQuantity] = useMutation(UPDATE_ORDER);

  const updateQuantity = () => {
    updateOrderQuantity({
      variables: {
        order: {
          id: order.id,
          quantity,
        },
      },
    });
  };

  return (
    <Grid2 size={6}>
      <Card variant="outlined">
        <CardContent>
          <Typography variant="h5" component="div">
            {order.name}
          </Typography>
          <Typography sx={{ color: "text.secondary", mb: 1.5 }}>
            {order.status}
          </Typography>
          <Typography variant="body2">{order.quantity}</Typography>
          <TextField
            id="outlined-basic"
            label="Quantity"
            variant="outlined"
            type="number"
            onChange={(e) => {
              setQuantity(Number(e.target.value));
            }}
            value={quantity}
          />
          <CardActions>
            <Button variant="contained" onClick={updateQuantity}>
              Update quantity
            </Button>
          </CardActions>
        </CardContent>
      </Card>
    </Grid2>
  );
};
