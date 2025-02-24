import React from "react";
import { View, Text, Button } from "react-native";
import { createStackNavigator } from "@react-navigation/stack";
import { NavigationContainer } from "@react-navigation/native";

import Home from './src/screens/Home'; // Import the Home component
import Login from './src/screens/Login'; // Import the Home component
import RegistrarEquipo from './src/screens/RegistrarEquipo'; // Import the Home component
import Marcadores from './src/screens/Marcadores'; // Import the Home component
import Competiciones from "./src/screens/Competiciones";
import CompeticionSumo from "./src/screens/CompeticionSumo";
import CompeticionSigueLineas from "./src/screens/CompeticionSigueLineas";


const Stack = createStackNavigator();
export default function App() {
  return (

    <NavigationContainer>
      <Stack.Navigator>
        <Stack.Screen name="Home" component={Home} />
        <Stack.Screen name="Marcadores" component={Marcadores} />
        <Stack.Screen name="RegistrarEquipo" component={RegistrarEquipo} />
        <Stack.Screen name="Login" component={Login} />
        <Stack.Screen name="Competiciones" component={Competiciones} />
        <Stack.Screen name="CompeticionSumo" component={CompeticionSumo} />
        <Stack.Screen name="CompeticionSigueLineas" component={CompeticionSigueLineas} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}

