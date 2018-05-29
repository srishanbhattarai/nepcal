class Nepcal < Formula
  desc "Equivalent of Linux's cal, for Nepali dates"
  homepage "https://github.com/nepcal/nepcal"
  url "https://github.com/nepcal/nepcal/archive/v0.2.0.tar.gz"
  head "https://github.com/nepcal/nepcal.git"
  version "0.2.0"
  sha256 "e00c2012069dc7e69f85000e8d8e33e7c6e3ab74d1de4cda4ac629ce54b6a9fa"
  depends_on "make" => :build

  def install
    bin.install "nepcal"
  end
end
