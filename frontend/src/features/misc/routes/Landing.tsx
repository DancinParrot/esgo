import { MainLayout } from "@/components/Layout";
import { BarChart } from "@mantine/charts";
import { Container, Grid, MultiSelect, Select, Stack } from "@mantine/core";
import { useEffect, useState } from "react";
import { getAllFields } from "../api/getAllFields";

const data = [
  { month: 'January', Smartphones: 120, Laptops: 80, Tablets: 10 },
  { month: 'February', Smartphones: 90, Laptops: 120, Tablets: 40 },
  { month: 'March', Smartphones: 40, Laptops: 100, Tablets: 20 },
  { month: 'April', Smartphones: 100, Laptops: 20, Tablets: 80 },
  { month: 'May', Smartphones: 80, Laptops: 140, Tablets: 120 },
  { month: 'June', Smartphones: 75, Laptops: 60, Tablets: 100 },
];

// Get columns from backend
// Get data from backend to show on table
// Based on column picked and type of chart, send request to aggregate on backend


export const Landing = () => {
  const [fields, setFields] = useState<string[]>([])
  const [dataKey, setDataKey] = useState<string | null>('');

  useEffect(() => {
    getAllFields()
      .then(res => {
        setFields(res.fields);
      })
      .catch(err => {
        console.error(err);
      })
  }, [])

  return (
    <MainLayout>
      <Container size={"xl"}>
        <Grid>
          <Grid.Col span={8}>
            <BarChart
              h={300}
              data={data}
              dataKey="month"
              series={[
                { name: 'Smartphones', color: 'violet.6' },
                { name: 'Laptops', color: 'blue.6' },
                { name: 'Tablets', color: 'teal.6' },
              ]}
              tickLine="y"
            />
          </Grid.Col>
          <Grid.Col span={4}>
            <Stack
              // h={300}
              bg="var(--mantine-color-body)"
            >
              <Select
                label="Type of chart"
                placeholder="Pick value"
                data={['Bar Chart', 'Pie Chart']}
              />
              <MultiSelect
                label="Column(s) to shown"
                placeholder="Pick value"
                data={fields}
              />
              <Select
                label="Data key"
                placeholder="Pick value"
                data={fields}
                onChange={setDataKey}
              />
            </Stack>
          </Grid.Col>
        </Grid>
      </Container>
    </MainLayout>
  );
};
