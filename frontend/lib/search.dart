import 'package:flutter/material.dart';
import 'ClassSearch.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'signup_login.dart';

class SearchPage extends StatelessWidget {
  const SearchPage({super.key}); // constを追加

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Search"),
        backgroundColor: Theme.of(context).colorScheme.primaryContainer,
      ),
      body: const SearchForm(),
    );
  }
}

class SearchForm extends StatefulWidget {
  const SearchForm({super.key});

  @override
  SearchFormState createState() => SearchFormState();
}

class SearchFormState extends State<SearchForm> {
  final TextEditingController _universityController = TextEditingController();
  final TextEditingController _fcultyController = TextEditingController();
  final TextEditingController _departmentController = TextEditingController();
  final TextEditingController _gradeController = TextEditingController();
  final TextEditingController _classController = TextEditingController();
  bool _isLoading = false;

  Future<void> _createclass(groupId) async {
    final String classname = _classController.text.trim();
    
    setState(() {
      _isLoading = true;
    });

    final url = Uri.parse('http://localhost:8080/classes');
    try {
      final response = await http.post(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ${TokenManager().accessToken}',
        },
        body: jsonEncode({
          'classname': classname,
          "group_id": groupId
        })
      );

      if (response.statusCode == 201) {
        if (mounted) {
          Navigator.pop(context);
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(content: Text('作成成功'))
          );
        }
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
          SnackBar(content: Text('Network error: $e'))
        );
      }
    } finally {
      if (mounted) {
        setState(() {
          _isLoading = false;
        });
      }
    }
  }

  Future<void> _creategroup() async {
    final String university = _universityController.text.trim();
    final String fculty = _fcultyController.text.trim();
    final String department = _departmentController.text.trim();
    final String grade = _gradeController.text.trim();

    setState(() {
      _isLoading = true;
    });

    final url = Uri.parse('http://localhost:8080/groups');
    try {
      final response = await http.post(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ${TokenManager().accessToken}',
        },
        body: jsonEncode({
          'university': university,
          'fculty': fculty,
          'department': department,
          'grade': grade,
        }),
      );

      if (response.statusCode == 201) {
        if (mounted) {
          Navigator.pop(context);
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(content: Text('作成成功'))
            );
        }
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
          SnackBar(content: Text('Network error: $e'))
        );
      }
    } finally {
      if (mounted) {
        setState(() {
          _isLoading = false;
        });
      }
    }
  }

  Future<void> _search() async {
    final String university = _universityController.text.trim();
    final String fculty = _fcultyController.text.trim();
    final String department = _departmentController.text.trim();
    final String grade = _gradeController.text.trim();
    final String classname = _classController.text.trim();

    setState(() {
      _isLoading = true;
    });

    final url = Uri.parse('http://localhost:8080/groups/get');
    try {
      final response = await http.post(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ${TokenManager().accessToken}',
          },
        body: jsonEncode({
          'university': university,
          'fculty': fculty,
          'department': department,
          'grade': grade,
        }),
      );

      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);
        final groupId = data['groups'][0]['ID']; // groupのIDを取得
        final classUrl = Uri.parse('http://localhost:8080/classes/get');
        try {
          final responseClass = await http.post(
            classUrl,
            headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer ${TokenManager().accessToken}',
            },
            body: jsonEncode({
              'classname': classname,
              'group_id': groupId,
            }
            ),
          );
          if (responseClass.statusCode == 200) {
            final classData = jsonDecode(responseClass.body);
            final classId = classData['classes'][0]['ID'];
            if (mounted) {
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (context) => ClassPage(classId: classId),
                ),
              );
            }
          } else {
            if (mounted) {
              await showDialog(
                context: context,
                builder: (BuildContext context) {
                  return SimpleDialog(
                    title: const Text('授業が存在しません。作成しますか？'),
                    children: [
                      SimpleDialogOption(
                        onPressed: () {
                          _createclass(groupId);
                        },
                        child: const Text('作成'),
                      ),
                      SimpleDialogOption(
                        onPressed: () {
                          Navigator.pop(context);
                        },
                        child: const Text('キャンセル'),
                      )
                    ],
                  );
                }
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
              _isLoading = false;
            });
          }
        }
      } else {
        if (mounted) {
          await showDialog(
            context: context,
            builder: (BuildContext context) {
              return SimpleDialog(
                title: const Text('グループが存在しません。作成しますか？'),
                children: [
                  SimpleDialogOption(
                    onPressed: _creategroup,
                    child: const Text('作成'),
                  ),
                  SimpleDialogOption(
                    onPressed: () {
                      Navigator.pop(context);
                    },
                    child: const Text('キャンセル'),
                  )
                ],
              );
            }
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
          _isLoading = false;
        });
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          TextField(
            controller: _universityController,
            decoration: const InputDecoration(
              labelText: 'University',
              border: OutlineInputBorder(),
            ),
          ),
          const SizedBox(height: 16),
          TextField(
            controller: _fcultyController,
            decoration: const InputDecoration(
              labelText: 'Faculty',
              border: OutlineInputBorder(),
            ),
          ),
          const SizedBox(height: 16),
          TextField(
            controller: _departmentController,
            decoration: const InputDecoration(
              labelText: 'Department',
              border: OutlineInputBorder(),
            ),
          ),
          const SizedBox(height: 16),
          TextField(
            controller: _gradeController,
            decoration: const InputDecoration(
              labelText: 'Grade',
              border: OutlineInputBorder(),
            ),
          ),
          const SizedBox(height: 16),
          TextField(
            controller: _classController,
            decoration: const InputDecoration(
              labelText: 'Class',
              border: OutlineInputBorder(),
            ),
          ),
          const SizedBox(height: 20),
          _isLoading
              ? const CircularProgressIndicator()
              : ElevatedButton(
                  onPressed: _search,
                  child: const Text("Search"),
                ),
        ],
      ),
    );
  }
}
