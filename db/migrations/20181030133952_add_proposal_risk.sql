
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO project_risks (project_id, risk, mitigation) VALUES
(1, 'Hujan', 'Menggunakan pupuk dan ZPT agar tanaman tahan terhadap cuaca.'),
(1, 'Angin Kencang', 'Menggunakan ajir dan tali yang kuat untuk menopang tanaman.'),
(1, 'Kekeringan', 'Menggunakan kolam tadah hujan, dan lokasi penanaman dekat mata air.'),
(1, 'Banjir', 'Tinggi minimal Bedengan 30 cm dan lahan terletak pada daerah perbukitan.'),
(1, 'Serangan Gulma', 'Kami menggunakan Mulsa Plastik untuk menekan gulma.'),
(1, 'Serangan Hama', 'Penggunaan pestisida pada tingkat serangan tertentu, menggunakan prinsip musuh alami dan tumpang sari.'),
(1, 'Harga Jatuh', 'Kami menggunakan proyeksi harga terendah di tingkat tengkulak, sehingga Analisis biaya yang kami tampilkan adalah kondisi terburuknya');

INSERT INTO project_risks (project_id, risk, mitigation) VALUES
(2, 'Hujan', 'Menggunakan pupuk dan ZPT agar tanaman tahan terhadap cuaca.'),
(2, 'Kekeringan', 'Menggunakan kolam tadah hujan, dan lokasi penanaman dekat mata air.'),
(2, 'Banjir', 'Tinggi minimal Bedengan 30 cm dan lahan terletak pada daerah perbukitan.'),
(2, 'Serangan Gulma', 'Pembersihan lahan dari gulma dilakukan rutin setiap sebulan 2 kali'),
(2, 'Serangan Hama', 'Penggunaan pestisida pada tingkat serangan tertentu, menggunakan prinsip musuh alami.'),
(2, 'Harga Jatuh', 'Kami menggunakan proyeksi harga terendah di tingkat tengkulak, sehingga Analisis biaya yang kami tampilkan adalah kondisi terburuknya.');

INSERT INTO project_risks (project_id, risk, mitigation) VALUES
(3, 'Hujan', 'Menggunakan pupuk dan ZPT agar tanaman tahan terhadap cuaca. Daun Bawang Tahan Hujan.'),
(3, 'Angin Kencang', 'Daun Bawang Tidak terlalu terpengaruh. Bulan Februari sudah bukan musim angin.'),
(3, 'Kekeringan', 'Lokasi penanaman dekat mata air.'),
(3, 'Banjir', 'Tinggi minimal Bedengan 25 cm dan lahan terletak pada daerah perbukitan.'),
(3, 'Serangan Gulma', 'Perawatan dan pengawasan Rutin setiap hari.'),
(3, 'Serangan Hama', 'Penggunaan pestisida pada tingkat serangan tertentu, menggunakan prinsip musuh alami dan tumpang sari.'),
(3, 'Harga Jatuh', 'Kami menggunakan proyeksi harga terendah di tingkat tengkulak, sehingga Analisis biaya yang kami tampilkan adalah kondisi terburuknya.');

INSERT INTO project_risks (project_id, risk, mitigation) VALUES
(4, 'Hujan', 'Menggunakan ajir dan tali yang kuat untuk menahan tanaman agar tidak roboh.'),
(4, 'Kekeringan', 'Lokasi penanaman dekat mata air dan sungai.'),
(4, 'Banjir', 'Lahan terletak pada daerah perbukitan.'),
(4, 'Serangan Gulma', 'Perawatan dan pengawasan Rutin setiap hari serta kegiatan penyiangan rutin setiap bulannya.'),
(4, 'Serangan Hama', 'Penggunaan pestisida pada tingkat serangan tertentu dapat juga menggunakan prinsip musuh alami dan pencegahan serangan hama secara mekanis.'),
(4, 'Harga Jatuh', 'Kami menggunakan proyeksi harga terendah di tingkat tengkulak, sehingga Analisis biaya yang kami tampilkan adalah kondisi terburuknya.');

INSERT INTO project_risks (project_id, risk, mitigation) VALUES
(5, 'Angin Kencang', 'Tanaman tidak terlalu tinggi dan kokoh. tidak terlalu berpengaruh terhadap angin.'),
(5, 'Kekeringan', 'Lokasi penanaman dekat mata air.'),
(5, 'Banjir', 'Lahan terletak pada daerah perbukitan.'),
(5, 'Serangan Gulma', 'Perawatan dan pengawasan Rutin setiap hari serta kegiatan penyiangan rutin setiap bulannya.'),
(5, 'Serangan Hama', 'Penggunaan pestisida pada tingkat serangan tertentu dapat juga menggunakan prinsip musuh alami dan pencegahan serangan hama secara mekanis.'),
(5, 'Harga Jatuh', 'Kami menggunakan proyeksi harga terendah di tingkat tengkulak, sehingga Analisis biaya yang kami tampilkan adalah kondisi terburuknya.');

INSERT INTO project_risks (project_id, risk, mitigation) VALUES
(6, 'Angin Kencang', 'Tanaman tidak terlalu tinggi dan kokoh. tidak terlalu berpengaruh terhadap angin.'),
(6, 'Kekeringan', 'Lokasi penanaman dekat mata air dan sungai.'),
(6, 'Banjir', 'Lahan terletak pada daerah perbukitan.'),
(6, 'Serangan Gulma', 'Perawatan dan pengawasan Rutin setiap hari serta kegiatan penyiangan rutin setiap bulannya.'),
(6, 'Serangan Hama', 'Penggunaan pestisida pada tingkat serangan tertentu dapat juga menggunakan prinsip musuh alami dan pencegahan serangan hama secara mekanis.'),
(6, 'Harga Jatuh', 'Kami menggunakan proyeksi harga terendah di tingkat tengkulak, sehingga Analisis biaya yang kami tampilkan adalah kondisi terburuknya.');

INSERT INTO project_risks (project_id, risk, mitigation) VALUES
(7, 'Hujan', 'Menggunakan pupuk dan ZPT agar tanaman tahan terhadap cuaca.'),
(7, 'Angin Kencang', 'Menggunakan ajir dan tali yang kuat untuk menopang tanaman.'),
(7, 'Kekeringan', 'Menggunakan kolam tadah hujan, dan lokasi penanaman dekat mata air.'),
(7, 'Banjir', 'Tinggi minimal Bedengan 30 cm dan lahan terletak pada daerah perbukitan.'),
(7, 'Serangan Gulma', 'Kami menggunakan Mulsa Plastik untuk menekan gulma.'),
(7, 'Serangan Hama', 'Penggunaan pestisida pada tingkat serangan tertentu, menggunakan prinsip musuh alami dan tumpang sari.'),
(7, 'Harga Jatuh', 'Kami menggunakan proyeksi harga terendah di tingkat tengkulak, sehingga Analisis biaya yang kami tampilkan adalah kondisi terburuknya.');

INSERT INTO project_risks (project_id, risk, mitigation) VALUES
(8, 'Kekeringan', 'Menggunakan kolam tadah hujan, lokasi penanaman dekat dengan waduk penampungan air.'),
(8, 'Banjir', 'Tinggi minimal Bedengan 40 cm dan lahan terletak pada daerah perbukitan.'),
(8, 'Serangan Gulma', 'Perawatan intensif oleh petani menggunakan alat pembumbun tanah sekaligus untuk mengangkat gulma.'),
(8, 'Serangan Hama', 'Penggunaan pestisida pada tingkat serangan tertentu, menggunakan prinsip musuh alami dan mekanisasi pertanian.'),
(8, 'Harga Jatuh', 'Kami menggunakan proyeksi harga terendah di tingkat tengkulak, sehingga Analisis biaya yang kami tampilkan adalah kondisi terburuknya.'),
(8, 'Kualitas Bibit', 'Mitra kami memiliki penangkaran bibit yang bekerjasama langsung dengan Balai Penelitian Tanaman Sayuran (Balitsa) Kementerian Pertanian.');

INSERT INTO project_risks (project_id, risk, mitigation) VALUES
(9, 'Kekeringan', 'Menggunakan kolam tadah hujan, lokasi penanaman dekat dengan waduk penampungan air.'),
(9, 'Banjir', 'Tinggi minimal Bedengan 40 cm dan lahan terletak pada daerah perbukitan.'),
(9, 'Serangan Gulma', 'Perawatan intensif oleh petani menggunakan alat pembumbun tanah sekaligus untuk mengangkat gulma.'),
(9, 'Serangan Hama', 'Penggunaan pestisida pada tingkat serangan tertentu, menggunakan prinsip musuh alami dan mekanisasi pertanian.'),
(9, 'Kualitas Bibit', 'Mitra kami memiliki penangkaran bibit yang bekerjasama langsung dengan Balai Penelitian Tanaman Sayuran (Balitsa) Kementerian Pertanian.'),
(9, 'Harga Jatuh', 'Kami menggunakan proyeksi harga terendah di tingkat tengkulak, sehingga Analisis biaya yang kami tampilkan adalah kondisi terburuknya.'),
(9, 'Kegagalan Panen', 'Peluang gagal panen amat kecil karena Tanijoy melakukan pilot project terlebih dahulu untuk menguji kinerja Petani Mitra dan hasil panen komoditas sebelum membuka proyek investasi.');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

