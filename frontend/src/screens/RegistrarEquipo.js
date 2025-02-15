import React, { useState } from 'react';
import { View, Text, TextInput, Button, StyleSheet } from 'react-native';

const RegistrarEquipo = () => {
    const [teamName, setTeamName] = useState('');
    const [teamMembers, setTeamMembers] = useState('');

    const handleRegister = () => {
        // Handle team registration logic here
        console.log('Team Name:', teamName);
        console.log('Team Members:', teamMembers);
    };

    return (
        <View style={styles.container}>
            <Text style={styles.title}>Register Team</Text>
            <TextInput
                style={styles.input}
                placeholder="Team Name"
                value={teamName}
                onChangeText={setTeamName}
            />
            <TextInput
                style={styles.input}
                placeholder="Team Members"
                value={teamMembers}
                onChangeText={setTeamMembers}
            />
            <Button title="Register" onPress={handleRegister} />
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        padding: 16,
    },
    title: {
        fontSize: 24,
        marginBottom: 16,
        textAlign: 'center',
    },
    input: {
        height: 40,
        borderColor: 'gray',
        borderWidth: 1,
        marginBottom: 12,
        paddingHorizontal: 8,
    },
});

export default RegistrarEquipo;