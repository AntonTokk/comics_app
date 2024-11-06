import 'package:flutter/material.dart';
import 'package:dio/dio.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Comics App',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: ComicsListScreen(),
    );
  }
}

class ComicsListScreen extends StatefulWidget {
  const ComicsListScreen({super.key});

  @override
  _ComicsListScreenState createState() => _ComicsListScreenState();
}

class _ComicsListScreenState extends State<ComicsListScreen> {
  List<dynamic> comics = [];
  bool isLoading = true;
  String errorMessage = '';
  final Dio _dio = Dio();

  @override
  void initState() {
    super.initState();
    fetchComics();
  }

  Future<void> fetchComics() async {
    try {
      final response = await _dio.get('http://10.0.2.2:8080/comics');
      if (response.statusCode == 200) {
        setState(() {
          comics = response.data;
          isLoading = false;
        });
      } else {
        setState(() {
          errorMessage = 'Failed to load comics';
          isLoading = false;
        });
      }
    } catch (e) {
      setState(() {
        errorMessage = 'Error fetching comics: $e';
        isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Comics'),
      ),
      body: isLoading
          ? const Center(child: CircularProgressIndicator())
          : errorMessage.isNotEmpty
              ? Center(child: Text(errorMessage))
              : GridView.builder(
                  gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                    crossAxisCount: 2, // Количество столбцов
                    childAspectRatio: 0.75, // Соотношение сторон элементов
                  ),
                  itemCount: comics.length,
                  itemBuilder: (context, index) {
                    return GestureDetector(
                      onTap: () {
                        Navigator.push(
                          context,
                          MaterialPageRoute(
                            builder: (context) => ComicDetailScreen(comic: comics[index]),
                          ),
                        );
                      },
                      child: Card(
                        child: Column(
                          children: [
                            Expanded(
                              child: Image.network(
                                comics[index]['imageUrl'],
                                fit: BoxFit.cover,
                              ),
                            ),
                            Padding(
                              padding: const EdgeInsets.all(8.0),
                              child: Text(
                                comics[index]['title'],
                                style: const TextStyle(fontWeight: FontWeight.bold),
                                overflow: TextOverflow.ellipsis,
                                maxLines: 2,
                              ),
                            ),
                            Padding(
                              padding: const EdgeInsets.all(8.0),
                              child: Text(
                                comics[index]['author'],
                                overflow: TextOverflow.ellipsis,
                                maxLines: 1,
                              ),
                            ),
                          ],
                        ),
                      ),
                    );
                  },
                ),
    );
  }
}

class ComicDetailScreen extends StatelessWidget {
  final dynamic comic;

  const ComicDetailScreen({super.key, required this.comic});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(comic['title']),
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Image.network(comic['imageUrl']),
              const SizedBox(height: 16),
              Text('Title: ${comic['title']}', style: const TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
              const SizedBox(height: 8),
              Text('Author: ${comic['author']}', style: const TextStyle(fontSize: 16)),
              const SizedBox(height: 8),
              Text('Description: ${comic['description']}', style: const TextStyle(fontSize: 16)),
              const SizedBox(height: 8),
              Text('Price: \$${comic['price']}', style: const TextStyle(fontSize: 16)),
              const SizedBox(height: 8),
              Text('Category: ${comic['category']}', style: const TextStyle(fontSize: 16)),
              const SizedBox(height: 8),
              Text('Quantity: ${comic['quantity']}', style: const TextStyle(fontSize: 16)),
              const SizedBox(height: 8),
              Text('Is Favorite: ${comic['isFavorite'] ? 'Yes' : 'No'}', style: const TextStyle(fontSize: 16)),
            ],
          ),
        ),
      ),
    );
  }
}