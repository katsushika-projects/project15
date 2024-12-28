import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class ClassSearchPage extends StatefulWidget {
  final String groupId;

  const ClassSearchPage({super.key, required this.groupId});

  @override
  ClassSearchPageState createState() => ClassSearchPageState();
}

class ClassSearchPageState extends State<ClassSearchPage> {
  List<dynamic> classes = [];
  bool isLoading = true;

  @override
  void initState() {
    super.initState();
    fetchClasses();
  }

  Future<void> fetchClasses() async {
    final url = Uri.parse('http://localhost:8080/classes?groupId=${widget.groupId}');
    try {
      final response = await http.get(url);

      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);
        setState(() {
          classes = data['classes']; // APIの結果からクラス一覧を取得
          isLoading = false;
        });
      } else {
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text('Server error: ${response.statusCode}')),
          );
        }
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Network error: $e')),
        );
      }
    } finally {
      if (mounted) {
        setState(() {
          isLoading = false;
        });
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Class Search Results"),
      ),
      body: isLoading
          ? const Center(child: CircularProgressIndicator())
          : classes.isEmpty
              ? const Center(child: Text("No classes found"))
              : ListView.builder(
                  itemCount: classes.length,
                  itemBuilder: (context, index) {
                    return ListTile(
                      title: Text(classes[index]['name']),
                      subtitle: Text('Class ID: ${classes[index]['id']}'),
                    );
                  },
                ),
    );
  }
}
